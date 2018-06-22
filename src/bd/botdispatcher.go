package bd

import (
	"crypto/ecdsa"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// BotDispatcher is a rpc server that allows the client to seamlessly dispatch bots
type BotDispatcher struct {

	// For adding new stakes
	bdLock sync.Mutex

	// A map from stakes to BotQueues
	queues map[float32]*BotQueues

	// Refilling key
	refillKey *ecdsa.PrivateKey

	// RPC stuff
	dead bool
	l    net.Listener
	port int
}

// BotQueues maintains the bots in three different queues: available, busy and refill
type BotQueues struct {

	// Lock for consistency
	qLock sync.Mutex

	// The bot ques
	available []*Bot
	//busy      []*Bot
	refill []*Bot
}

// Bot is an interface to a bot. A bot must be able to do bot stuff
type Bot interface {
	DoBotStuff() error
}

// Init initializes the BotDispatcher on ipAddr:port.
// hexkey is the key to unlock the funding wallet
func (bd *BotDispatcher) Init(ipAddr string, port int, hexkey string) error {

	bd.bdLock.Lock()
	defer bd.bdLock.Lock()

	// Convert the string into a ecdsa privkey
	privk, err := ethcrypto.HexToECDSA(hexkey)
	if err != nil {
		log.Fatalf("GM: bad private key")
	}
	bd.refillKey = privk

	// gm.auth = bind.NewKeyedTransactor(privk)
	// if gm.auth == nil {
	// 	log.Fatalf("GM: failed to create authorized transactor: %v", err)
	// }

	bd.queues = make(map[float32]*BotQueues)

	// RPC RELATED STUFF BELOW
	// Register our baby with net/rpc
	bd.port = port

	rpcs := rpc.NewServer()
	rpcs.Register(bd)

	// Create a TCP listener that will listen on `Port`
	l, e := net.Listen("tcp", ipAddr+":"+strconv.Itoa(bd.port))
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
	log.Printf("INFO GM: Initialization succesful.\n")

	return nil
}

// Kill stops the BotDispatcher
func (bd *BotDispatcher) Kill() {

	bd.bdLock.Lock()
	defer bd.bdLock.Lock()

	bd.dead = true
	bd.l.Close()

	log.Printf("INFO BotDispatcher: dead\n")

}
