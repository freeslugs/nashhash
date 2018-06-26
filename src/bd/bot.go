package bd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"log"
	"strconv"
	"time"

	gorand "math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/crypto"
)

// Bot is a automated agent. Bot will DoBotStuff on a contract
type Bot struct {
	key  *ecdsa.PrivateKey
	auth *bind.TransactOpts
}

// Init the Bot. Create a public private key pair at minimum
func (b *Bot) Init() error {

	// Create if
	sk, e := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if e != nil {
		return e
	}
	b.key = sk

	// Create the authenticator
	auth := bind.NewKeyedTransactor(sk)
	if auth == nil {
		return errors.New("failed to create new transactor")
	}

	b.auth = auth

	log.Printf("KEY: %x\n", crypto.FromECDSA(sk))

	return nil
}

// Kill the Bot and harvest its balance to the address
func (b *Bot) Kill(address common.Address) error {

	// Lets send the remaining balance back
	e := harvestAccount(b.key, address)
	if e != nil {
		return e
	}

	return nil
}

// Balance returns the balance of the Bot
func (b *Bot) Balance() (float64, error) {

	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return -1.0, err
	}
	defer conn.Close()

	money, err := conn.BalanceAt(context.Background(), crypto.PubkeyToAddress(b.key.PublicKey), nil)
	if err != nil {
		return -1.0, err
	}

	return toEth(money), nil

}

// PendingBalance returns the pending balance of the bot
func (b *Bot) PendingBalance() (float64, error) {

	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return -1.0, err
	}
	defer conn.Close()

	money, err := conn.PendingBalanceAt(context.Background(), crypto.PubkeyToAddress(b.key.PublicKey))
	if err != nil {
		return -1.0, err
	}

	return toEth(money), nil

}

// DoBotStuff is a stateless function that performs automated
// behaviour on a contract at address
func (b *Bot) DoBotStuff(address string) error {

	log.Println("beep beep... doing bots stuff")

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(address), conn)
	if err != nil {
		return err
	}

	stake, err := game.GetStakeSize(nil)
	if err != nil {
		return err
	}

	s1 := gorand.NewSource(time.Now().UnixNano())
	r1 := gorand.New(s1)

	// Step 1: Commit all the guesses

	// Get our guess and secret, and get the hash.
	guessstr := strconv.Itoa(r1.Intn(60))
	secretstr := strconv.Itoa(r1.Intn(100000))
	guess := []byte(guessstr)
	secret := []byte(secretstr)
	h := crypto.Keccak256(guess, secret)
	var hash [32]byte
	copy(hash[:], h[:32])

	// Prep the auth and commit
	b.auth.Value = stake
	tx, txerr := game.Commit(b.auth, hash)
	if txerr != nil {
		log.Printf(txerr.Error())
		return txerr
	} else {
		log.Printf("bot commit succesful 0x%x\n", tx.Hash())
	}
	b.auth.Value = nil

	log.Printf("INFO Bot.DoBotStuff 0x%x: succesful bot commit\n", crypto.PubkeyToAddress(b.key.PublicKey))

	// We wait for the contract change state

	// Step 2: Wait for contract to be in reveal state
	inCommit := true
	for inCommit {
		state, err := game.GetGameStateInfo(nil)
		if err != nil {
			return err
		}

		switch s := state.State.Int64(); s {
		case 0:
			time.Sleep(5 * time.Second)
			continue
		case 1:
			inCommit = false
		default:
			return errors.New("bad contract state")
		}
	}

	// Step 3: Reveal all the guesses
	tx, txerr = game.Reveal(b.auth, guessstr, secretstr)
	if txerr != nil {
		return txerr
	}
	log.Printf("INFO Bot.DoBotStuff 0x%x: succesful bot reveal succesful\n", crypto.PubkeyToAddress(b.key.PublicKey))

	return nil
}
