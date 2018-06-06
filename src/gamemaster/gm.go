package gm

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

// The GameMaster object. Runs on the server. Manages GameOperators, connects
// new games, resets games and etc etc etc. Important mister.
type GameMaster struct {
	//handlers []GameOperator
	dead bool
	l    net.Listener
	port int
}

// Init initializes the game master. In particular, it should register the
// game master for RPC.
func (gm *GameMaster) Init(ipAddr string, port int) error {

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
	fmt.Println("gm is waiting...")

	return nil
}

// Execute is a test
func (gm *GameMaster) Execute(req ExecuteCallArgs, res *ExecuteCallReply) (err error) {
	if req.Message == "" {
		err = errors.New("You must give me a message")
		return
	}

	res.Response = "This is your message: " + req.Message
	return
}

// Kill the gamemaster is something is wrong.
func (gm *GameMaster) Kill() {
	gm.dead = true
	gm.l.Close()
}
