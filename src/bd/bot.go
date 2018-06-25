package bd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

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

	log.Printf("KEY: %x\n", crypto.FromECDSA(sk))

	return nil
}

// Kill the Bot and harvest its balance to the address
func (b *Bot) Kill(address string) error {

	// Lets send the remaining balance back
	e := harvestAccount(b.key, common.HexToAddress(address))
	if e != nil {
		return e
	}

	return nil
}

// DoBotStuff is a stateless function that performs automated
// behaviour on a contract at address
func (b *Bot) DoBotStuff(address string) error {

	log.Printf("beep bop bot doing bot stuff %s\n", address)
	return nil
}
