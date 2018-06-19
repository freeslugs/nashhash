package gm

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"strconv"

	"fmt"
	"log"
	"math/big"
	gorand "math/rand"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BotDispatcher controls operation of bots
type BotDispatcher struct {
	stakeSize       *big.Int // The stake that bot needs to commit
	contractAddress string   // The address of the party

	botPoolSize     int                  // how many bots we have at the party
	bdLock          sync.Mutex           // probaly unnecessary
	bots            []*bind.TransactOpts // slice of bots that are used to bet
	botKeys         []*ecdsa.PrivateKey
	sugarBot        *bind.TransactOpts // address that sponsors the party
	sugarBotKey     *ecdsa.PrivateKey
	sugarBotKeyLock *sync.Mutex

	refilldead chan bool
}

// Init initializes the dispatcher by creating the bot addresses and
// providing initial financing.
func (bd *BotDispatcher) Init(botPoolSize int, contractAddress string, sugarBot *bind.TransactOpts,
	sugarBotKey *ecdsa.PrivateKey, sugarBotKeyLock *sync.Mutex) error {

	bd.botPoolSize = botPoolSize
	bd.contractAddress = contractAddress
	bd.sugarBot = sugarBot
	bd.sugarBotKey = sugarBotKey
	bd.sugarBotKeyLock = sugarBotKeyLock
	bd.refilldead = make(chan bool)

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(bd.contractAddress), conn)
	if err != nil {
		return err
	}

	stake, err := game.GetStakeSize(nil)
	if err != nil {
		return err
	}

	bd.stakeSize = stake

	// Generate the private keys
	pubkeyCurve := crypto.S256()
	for i := 0; i < bd.botPoolSize; i++ {

		// Generate a private key
		privk, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
		if err != nil {
			log.Printf("ERROR BotDispatcher %s: %s\n", bd.contractAddress, err.Error())
			return err
		}
		bd.botKeys = append(bd.botKeys, privk)

		// Create the authenticator
		auth := bind.NewKeyedTransactor(privk)
		if auth == nil {
			log.Fatalf("ERROR BotDispatcher %s: %s\n", bd.contractAddress, err.Error())
		}
		bd.bots = append(bd.bots, auth)

	}

	f, err := os.OpenFile(BotKeysFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// Before we move any money into the bots, we should first record the keys
	for _, sk := range bd.botKeys {
		skstr := fmt.Sprintf("%x\n", crypto.FromECDSA(sk))
		f.WriteString(skstr)
	}

	// Fill up the wallet adresses
	go bd.refill()

	log.Printf("INFO BotDispatcher %s: init succesful\n", bd.contractAddress)
	return nil
}

// Kill destroys the bot dispatcher, returns all the funds to the sugarBot
func (bd *BotDispatcher) Kill() error {

	// Kill the refill routine
	bd.refilldead <- true

	// Route the remaining money back to the sponsor wallet
	err := harvestAccounts(bd.botKeys, crypto.PubkeyToAddress(bd.sugarBotKey.PublicKey))
	if err != nil {
		return err
	}

	log.Printf("INFO BotDispatcher %s: kill succesful\n", bd.contractAddress)
	return nil
}

// Guess stores the guess and the secret
type guess struct {
	guessstr  string
	secretstr string
}

// Dispatch sends the bots in. If how many is negative, we let the dispatcher
// decide how many bots to send. Dispatch is usually executed as a go routine, as
// it might take some time.
func (bd *BotDispatcher) Dispatch(howMany int) error {

	// We want to make sure
	if howMany > bd.botPoolSize {
		log.Printf("WARNING BotDispatcher %s: bot pool not big enough\n", bd.contractAddress)
		howMany = bd.botPoolSize
	}

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(bd.contractAddress), conn)
	if err != nil {
		return err
	}

	s1 := gorand.NewSource(time.Now().UnixNano())
	r1 := gorand.New(s1)

	// Step 1: Commit all the guesses
	var guesses []guess
	for i := 0; i < howMany; i++ {

		// Get our guess and secret, and get the hash.
		guessstr := strconv.Itoa(r1.Intn(60))
		secretstr := strconv.Itoa(r1.Intn(100000))
		guess := []byte(guessstr)
		secret := []byte(secretstr)
		h := crypto.Keccak256(guess, secret)
		var hash [32]byte
		copy(hash[:], h[:32])

		// Prep the auth and commit
		bd.bots[i].Value = bd.stakeSize
		tx, txerr := game.Commit(bd.bots[i], hash)
		if txerr != nil {
			log.Printf(txerr.Error())
		} else {
			log.Printf("bot commit succesful 0x%x\n", tx.Hash())
		}
		bd.bots[i].Value = nil

		guesses[i].guessstr = guessstr
		guesses[i].secretstr = secretstr
	}

	// Step 2: Wait for contract to be in reveal state
	inCommit := true
	for inCommit {
		state, err := game.GetGameStateInfo(nil)
		if err != nil {
			return err
		}

		switch s := state.State.Int64(); s {
		case GameCommitState:
			time.Sleep(2 * time.Second)
			continue
		case GameRevealState:
			inCommit = false
		default:
			return errors.New("bad contract state")
		}
	}

	// Step 3: Reveal all the guesses
	for i := 0; i < howMany; i++ {

		guessstr := guesses[i].guessstr
		secretstr := guesses[i].secretstr

		tx, txerr := game.Reveal(bd.bots[i], guessstr, secretstr)
		if txerr != nil {
			return txerr
		}
		log.Printf("bot reveal succesful 0x%x\n", tx.Hash())

	}

	return nil
}

// refill checks if any of the bots need a refill. This is a helper observer,
// that is launched from Init() and is killed in Kill(). Runs in a separate go routine.
func (bd *BotDispatcher) refill() {

	log.Printf("INFO BotDispatcher %s: refill routine started\n", bd.contractAddress)
	var limit big.Int
	limit.Mul(bd.stakeSize, big.NewInt(MinimumBalanceInStake))
	var refillAmount big.Int
	refillAmount.Mul(bd.stakeSize, big.NewInt(RefillAmountInStake))

	for i := 0; i < bd.botPoolSize; i = (i + 1) % bd.botPoolSize {

		select {
		// If we receive a kill signal, stop the goroutine
		case <-bd.refilldead:
			log.Printf("INFO BotDispatcher %s: refill routine terminated\n", bd.contractAddress)
			return

		// Otherwise, we refill the address
		default:

			// Get the balance of one of the bots
			conn, err := ethclient.Dial(EthClientPath)
			if err != nil {
				log.Printf("Failed to connect to the Ethereum client: %v", err)
				continue
			}
			money, err := conn.PendingBalanceAt(context.Background(), bd.bots[i].From)
			if err != nil {
				log.Printf("ERROR BotDispatcher %s: failed to get bot balance %s\n", bd.contractAddress, err.Error())
				conn.Close()
				continue
			}

			// If the balance is less than a certain minimum, we refiill
			// Not a one step thing....
			if money.Cmp(&limit) < 0 {

				err := sendEthSafe(
					bd.sugarBotKey,
					bd.bots[i].From,
					&refillAmount,
					bd.sugarBotKeyLock)

				if err != nil {
					log.Printf("sendEth failed %s\n", err.Error())
					conn.Close()
					continue
				}

				log.Printf("INFO BotDispatcher %s: succesful refill of %s\n", bd.contractAddress, bd.bots[i].From.Hex())

			}

			conn.Close()
			time.Sleep(1 * time.Second)
		}

	}
}
