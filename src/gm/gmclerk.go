package gm

import (
	"log"
	"net/rpc"
)

// Clerk is client side interface for the GameMaster.
// Still debating if this does more harm than good.
type Clerk struct {
	GMAddr string
	c      *rpc.Client
}

// Init initializes the Clerk to a GN address
func (gmk *Clerk) Init(srv string) error {
	gmk.GMAddr = srv
	c, err := rpc.Dial("tcp", srv)
	if err != nil {
		return err
	}

	gmk.c = c
	log.Printf("Clerk: initialization succesful\n")

	return nil
}

// Kill releases the resources held by the clerk.
func (gmk *Clerk) Kill() {
	gmk.c.Close()
}

// ConnectGame connects a game at @gameAddr to the gamemaster.
// By default, GM will start managing the states of the game
// after succesful connection
func (gmk *Clerk) ConnectGame(gameAddr string) error {
	e := connectGame(gmk.GMAddr, gameAddr)
	return e
}

// DisconnectGame disconnect a game at @gameAddr from the master
func (gmk *Clerk) DisconnectGame(gameAddr string) error {
	e := disconnectGame(gmk.GMAddr, gameAddr)
	return e
}
