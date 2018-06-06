package gm

import (
	"errors"
	"net"
)

// The GameMaster object. Runs on the server. Manages GameOperators, connects
// new games, resets games and etc etc etc. Important mister.
type GameMaster struct {
	//handlers []GameOperator
	dead bool
	l    net.Listener
}

// Init initializes the game master. In particular, it should register the
// game master for RPC.
func (gm *GameMaster) Init() error {
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
