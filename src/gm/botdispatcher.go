package gm

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// BotDispatcher controls operation of bots
type BotDispatcher struct {
	stakeSize       *big.Int // The stake that bot needs to commit
	contractAddress string   // The address of the party

	botPoolSize int                  // how many bots we have at the party
	bdLock      sync.Mutex           // probaly unnecessary
	bots        []*bind.TransactOpts // slice of bots that are used to bet
	sugarBot    *bind.TransactOpts   // address that sponsors the party
}

// Init initializes the dispatcher by creating the bot addresses and
// providing initial financing.
func (bd *BotDispatcher) Init(botPoolSize int) error {
	return nil
}

// Kill destroys the bot dispatcher, returns all the funds to the sugarBot
func (bd *BotDispatcher) Kill() error {
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
func (bd *BotDispatcher) refill() error {
	return nil
}
