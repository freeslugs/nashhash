package gm

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BotDispatcher controls operation of bots
type BotDispatcher struct {
	stakeSize       *big.Int // The stake that bot needs to commit
	contractAddress string   // The address of the party

	botPoolSize int                  // how many bots we have at the party
	bdLock      sync.Mutex           // probaly unnecessary
	bots        []*bind.TransactOpts // slice of bots that are used to bet
	botKeys     []*ecdsa.PrivateKey
	sugarBot    *bind.TransactOpts // address that sponsors the party
	sugarBotKey *ecdsa.PrivateKey

	refilldead chan bool
}

// Init initializes the dispatcher by creating the bot addresses and
// providing initial financing.
func (bd *BotDispatcher) Init(botPoolSize int, contractAddress string, sugarBot *bind.TransactOpts) error {

	bd.botPoolSize = botPoolSize
	bd.contractAddress = contractAddress
	bd.sugarBot = sugarBot

	// Generate the private keys
	pubkeyCurve := elliptic.P256()
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

	// Before we move any money iinto the bots, we should first record the keys
	for _, sk := range bd.botKeys {
		log.Println(sk)
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

	log.Printf("INFO BotDispatcher %s: kill succesful\n", bd.contractAddress)
	return nil
}

// Dispatch sends the bots in. If how many is negative, we let the dispatcher
// decide how many bots to send. Dispatch is usually executed as a go routine, as
// it might take some time.
func (bd *BotDispatcher) Dispatch(howMany int) error {
	return nil
}

// refill checks if any of the bots need a refill. This is a helper observer,
// that is launched from Init() and is killed in Kill(). Runs in a separate go routine.
func (bd *BotDispatcher) refill() {

	var nonce uint64
	var limit big.Int
	limit.Mul(bd.stakeSize, big.NewInt(5))
	var refillAmount big.Int
	refillAmount.Mul(&limit, big.NewInt(3))

	for i := 0; i < bd.botPoolSize; i++ {
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
			}
			money, err := conn.BalanceAt(context.Background(), bd.bots[i].From, nil)
			if err != nil {
				log.Printf("ERROR BotDispatcher %s: failed to get bot balance %s\n", bd.contractAddress, err.Error())
			}

			// If the balance is less than a certain minimum, we refiill
			// Not a one step thing....
			if money.Cmp(&limit) < 0 {

				// We need to ask the client about currect gas price
				gasPrice, err := conn.SuggestGasPrice(context.Background())
				if err != nil {
					log.Printf("ERROR BotDispatcher %s: gas price estimation failed\n", bd.contractAddress)
					continue
				}

				// This is the transaction to move money
				tx := types.NewTransaction(
					nonce,
					bd.bots[i].From,
					&refillAmount,
					21000, gasPrice, nil)

				signature, _ := crypto.Sign(tx.Hash().Bytes(), bd.sugarBotKey)
				signedtx, _ := tx.WithSignature(types.NewEIP155Signer(nil), signature)
				conn.SendTransaction(context.Background(), signedtx)
			}

		}

		// Doing iteration this way to check the dead channel often
		i++
		i = i % bd.botPoolSize
		nonce++
	}
}
