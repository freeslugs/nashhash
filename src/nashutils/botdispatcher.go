// Package bd gives you what you need to dispatch bots to your contracts.
package nashutils

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"net"
	"net/rpc"
	"strconv"
	"sync"
	"time"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// BotDispatcher is a rpc server that allows the client to seamlessly dispatch bots
type BotDispatcher struct {

	// For adding new stakes
	bdLock sync.Mutex // might have to change this r/w lock

	// A map from guaranteed balance and the different BotQ
	queues map[float64]*BotQ

	// Refilling
	refillKey  *ecdsa.PrivateKey
	refilldead chan bool

	// RPC stuff
	dead bool
	l    net.Listener
	port int
	ip   string
}

// Dispatch asks the bot dispatcher to dispatch args.Number bots to do Bot.BotStuff at address
// args.ContractAddress. If the DoBotStuff involves payable functions, you need to provide the
// balance the bots are expected to have in args.BotAllowance
func (bd *BotDispatcher) Dispatch(args DispatchArgs, res *DispatchReply) error {

	// Step 1: Find the appropriate BotQ. If we cannot gurantee this particular balance, retur
	bal := bd.findRightBalance(args.RequiredBalance)
	if bal == 0 {
		e := fmt.Errorf("required balance of %f ETH cannot be guaranteed",
			args.RequiredBalance)
		return e
	}
	botq := bd.queues[bal]

	// Step 2: Now that we have the correct BotQ we shall ask it to dispatch the bots
	e := botq.Dispatch(args.Number, args.ContractAddress)
	if e != nil {
		return e
	}

	log.Printf("INFO BotDispatcher.Dispatch: dispatching %d bots to %s, balance >= %f\n",
		args.Number, args.ContractAddress, args.RequiredBalance)
	return nil
}

// Init initializes the BotDispatcher on ipAddr:port.
// hexkey is the key to unlock the funding wallet
func (bd *BotDispatcher) Init(ipAddr string, port int, hexkey string, def bool) error {

	bd.bdLock.Lock()
	defer bd.bdLock.Unlock()

	// Convert the string into a ecdsa privkey
	privk, err := ethcrypto.HexToECDSA(hexkey)
	if err != nil {
		log.Fatalf("GM: bad private key")
	}
	bd.refillKey = privk

	// Lets create the que
	bd.queues = make(map[float64]*BotQ)
	if def {
		bd.initBotQsDefault()
	}

	// RPC stuff
	bd.port = port
	bd.ip = ipAddr
	bd.initRPC()

	// Start the refill routine
	bd.refilldead = make(chan bool)
	go bd.refill()

	return nil
}

// Kill stops the BotDispatcher
func (bd *BotDispatcher) Kill() {

	bd.bdLock.Lock()
	defer bd.bdLock.Unlock()

	// RPC stuff
	bd.dead = true
	bd.l.Close()

	// Terminate the refill routine first.
	bd.refilldead <- true

	for _, bq := range bd.queues {
		bq.Kill(ethcrypto.PubkeyToAddress(bd.refillKey.PublicKey))
	}

	log.Printf("INFO BotDispatcher: dead\n")

}

// helpers
//

// findRightalance will return the lowest supported guaranteed balance that is
// less than needed balance
func (bd *BotDispatcher) findRightBalance(balanceNeeded float64) float64 {

	lowestRightKey := math.MaxFloat64

	// We need to find a smallest key that is more than the balanceNeeded
	for k := range bd.queues {

		// If the key is smaller, that key does not work for us
		if k < balanceNeeded {
			continue
		}

		if k < lowestRightKey {
			lowestRightKey = k
		}
	}

	// If no guaranteed balance was bigger than the required balance, we return 0
	if lowestRightKey == math.MaxFloat64 {
		return 0
	}

	return lowestRightKey
}

func (bd *BotDispatcher) refill() {

	refillSleepTime := time.Duration(10)

	// We wait for commands on the channel.
	log.Println("INFO BotDispatcher: refill routine started started")
	for {
		select {
		case <-bd.refilldead:
			return
		default:

			// We need to iterate over all the qs
			for _, bq := range bd.queues {
				bq.Refill(bd.refillKey)
			}

			time.Sleep(refillSleepTime * time.Second)
		}
	}

}

// init helpers
//

func (bd *BotDispatcher) initBotQsDefault() error {

	amounts := [...]float64{0.05}
	defaultBotNumber := uint(5)

	for _, amount := range amounts {

		// Create the BotQ
		botq := &BotQ{}
		e := botq.Init(amount, defaultBotNumber)
		if e != nil {
			return e
		}
		bd.queues[amount] = botq

	}

	return nil
}

func (bd *BotDispatcher) initRPC() {

	rpcs := rpc.NewServer()
	rpcs.Register(bd)

	// Create a TCP listener that will listen on `Port`
	l, e := net.Listen("tcp", bd.ip+":"+strconv.Itoa(bd.port))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	bd.l = l
	// Go routine that accepts and serves new procedure calls
	go func() {
		for bd.dead == false {
			conn, err := bd.l.Accept()
			if err == nil && bd.dead == false {
				go rpcs.ServeConn(conn)
			} else if err == nil {
				conn.Close()
			}
			if err != nil && bd.dead == false {
				log.Printf("ERROR BotDispatcher accept: %v\n", err.Error())
				bd.Kill()
			}
		}
	}()
	log.Printf("INFO BotDispatcher: Initialization succesful.\n")

}
