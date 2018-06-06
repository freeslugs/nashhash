package gm

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
)

// The GameMaster object. Runs on the server. Manages GameOperators, connects
// new games, resets games and etc etc etc. Important mister.
type GameMaster struct {
	// handlers      []GameOperator
	operatedGames map[string]*GameOperator
	gmLock        sync.Mutex

	// RPC stuff
	dead bool
	l    net.Listener
	port int
}

// Init initializes the game master. In particular, it should register the
// game master for RPC.
func (gm *GameMaster) Init(ipAddr string, port int) error {

	gm.operatedGames = make(map[string]*GameOperator)

	// RPC RELATED STUFF BELOW
	// Register our baby with net/rpc
	gm.port = port

	rpcs := rpc.NewServer()
	rpcs.Register(gm)

	// Create a TCP listener that will listen on `Port`
	l, e := net.Listen("tcp", ipAddr+":"+strconv.Itoa(gm.port))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	gm.l = l

	// Go routine that accepts and serves new procedure calls
	go func() {
		for gm.dead == false {
			conn, err := gm.l.Accept()
			if err == nil && gm.dead == false {
				go rpcs.ServeConn(conn)
			} else if err == nil {
				conn.Close()
			}
			if err != nil && gm.dead == false {
				fmt.Printf("GameMaster accept: %v\n", err.Error())
				gm.Kill()
			}
		}
	}()
	fmt.Println("gm initialization succesful...")

	return nil
}

// Execute is a test
func (gm *GameMaster) Execute(req ExecuteCallArgs, res *ExecuteCallReply) error {
	if req.Message == "" {
		return errors.New("You must give me a message")
	}

	res.Response = "This is your message: " + req.Message
	return nil
}

// Connect call connects a GameOperator to a game at ConnectCallArgs.ContractAddress
func (gm *GameMaster) Connect(args ConnectCallArgs, res *ConnectCallReply) error {
	gm.gmLock.Lock()
	defer gm.gmLock.Unlock()

	// First we check if the game already has an operator on it
	addr := args.ContractAddress
	if gm.isOperated(addr) {
		return errors.New("GameMaster: game already operated")
	}

	// Create a game operator
	gop := &GameOperator{}
	gop.Init(addr)

	// Add this GameOperator to the mapping
	gm.operatedGames[addr] = gop

	// TODO: Give an option to not immediately start operating the game
	e := gm.operatedGames[addr].Play()
	if e != nil {
		panic("Error: inconsistent state in GM.Connect()")
		//		return e
	}

	return nil
}

// Disconnect call disconnects a GameOperator from a game at ConnectCallArgs.ContractAddress
func (gm *GameMaster) Disconnect(args DisconnectCallArgs, res *DisconnectCallReply) error {
	gm.gmLock.Lock()
	defer gm.gmLock.Unlock()

	// First we check if the game already has an operator on it
	addr := args.ContractAddress
	if !gm.isOperated(addr) {
		return errors.New("GameMaster: game at " + addr + " not operated")
	}

	// Stop the handler
	e := gm.operatedGames[addr].Stop()
	if e != nil {
		panic("Error: inconsistent state in GM.Connect()")
		//		return e
	}

	// Remove the gameOperator from the map
	delete(gm.operatedGames, addr)

	return nil
}

// Helper, checks if a game is already operated
func (gm *GameMaster) isOperated(addr string) bool {
	if _, ok := gm.operatedGames[addr]; ok {
		return true
	}
	return false
}

// Kill the gamemaster is something is wrong.
func (gm *GameMaster) Kill() {
	gm.dead = true
	gm.l.Close()
}
