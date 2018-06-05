package gm

import (
	"errors"
	"fmt"
)

const (
	StopGame = 0
)

// GameOperator operates a game contract
type GameOperator struct {
	contractAddress string
	controlChannel  chan int
	playing         bool
}

// ContractAddress returns the address of the game contract
// the handler is operating
func (gop *GameOperator) ContractAddress() string {
	return gop.contractAddress
}

// Playing returns true of the operator is playing a game
// right now
func (gop *GameOperator) Playing() bool {
	return gop.playing
}

// Init the handler
func (gop *GameOperator) Init(addr string) {
	gop.contractAddress = addr
	gop.controlChannel = make(chan int)
	gop.playing = false
}

// Play the game at the contract address
func (gop *GameOperator) Play() error {
	if gop.playing == true {
		return errors.New("")
	}
	go gop.playGame()
	return nil
}

func (gop *GameOperator) playGame() {
	fmt.Printf("playing the game on address %s\n", gop.contractAddress)
	<-gop.controlChannel
	fmt.Printf("operator quitting the game at address %s\n", gop.contractAddress)
}

// Stop stops the operator from operating the game. The game ideally should also
// be paused.
func (gop *GameOperator) Stop() error {
	if gop.playing == false {
		return errors.New("game already stopped")
	}
	gop.controlChannel <- StopGame
	return nil
}
