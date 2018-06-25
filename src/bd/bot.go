package bd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"log"

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

	// // We want to make sure
	// if howMany > bd.botPoolSize {
	// 	log.Printf("WARNING BotDispatcher %s: bot pool not big enough\n", bd.contractAddress)
	// 	howMany = bd.botPoolSize
	// }

	// // Create an IPC based RPC connection to a remote node
	// conn, err := ethclient.Dial(EthClientPath)
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close()

	// // Instantiate the contract and display its name
	// game, err := NewGame(common.HexToAddress(bd.contractAddress), conn)
	// if err != nil {
	// 	return err
	// }

	// round, err := game.GetRound(nil)
	// if err != nil {
	// 	return err
	// }

	// // We can only play if the last played round is strictly less then the current
	// // game round
	// if bd.round >= round.Int64() {
	// 	return nil
	// }

	// bd.round = round.Int64()

	// s1 := gorand.NewSource(time.Now().UnixNano())
	// r1 := gorand.New(s1)

	// // Step 1: Commit all the guesses
	// crdata := make([]CRData, howMany)
	// for i := 0; i < howMany; i++ {

	// 	// Get our guess and secret, and get the hash.
	// 	guessstr := strconv.Itoa(r1.Intn(60))
	// 	secretstr := strconv.Itoa(r1.Intn(100000))
	// 	guess := []byte(guessstr)
	// 	secret := []byte(secretstr)
	// 	h := crypto.Keccak256(guess, secret)
	// 	var hash [32]byte
	// 	copy(hash[:], h[:32])

	// 	// Prep the auth and commit
	// 	bd.bots[i].Value = bd.stakeSize
	// 	tx, txerr := game.Commit(bd.bots[i], hash)
	// 	if txerr != nil {
	// 		log.Printf(txerr.Error())
	// 	} else {
	// 		log.Printf("bot commit succesful 0x%x\n", tx.Hash())
	// 	}
	// 	bd.bots[i].Value = nil

	// 	crdata[i].Guess = guessstr
	// 	crdata[i].Secret = secretstr

	// 	log.Printf("INFO BotDispatcher.Dispatch(): succesful bot commit %d\n", i)
	// }

	// // Step 2: Wait for contract to be in reveal state
	// inCommit := true
	// for inCommit {
	// 	state, err := game.GetGameStateInfo(nil)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	switch s := state.State.Int64(); s {
	// 	case GameCommitState:
	// 		time.Sleep(2 * time.Second)
	// 		continue
	// 	case GameRevealState:
	// 		inCommit = false
	// 	default:
	// 		return errors.New("bad contract state")
	// 	}
	// }

	// // Step 3: Reveal all the guesses
	// for i := 0; i < howMany; i++ {

	// 	guessstr := crdata[i].Guess
	// 	secretstr := crdata[i].Secret

	// 	tx, txerr := game.Reveal(bd.bots[i], guessstr, secretstr)
	// 	if txerr != nil {
	// 		return txerr
	// 	}
	// 	log.Printf("bot reveal succesful 0x%x\n", tx.Hash())

	// }

	return nil
}
