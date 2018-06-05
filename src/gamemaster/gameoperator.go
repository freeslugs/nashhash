package gm

import (
	"errors"
	"fmt"
	"time"
)

const (
	// StopGame command to stop a game
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
	gop.playing = true
	return nil
}

// Game control logic goes into this function
func (gop *GameOperator) playGame() {
	fmt.Printf("playing the game on address %s\n", gop.contractAddress)
	for {

		// We use select here for non-blocking read on the socket
		select {
		case cmd := <-gop.controlChannel:

			// Different commands can be passed here.
			// For example, we can ask the handler to reset the state
			// or log the current state of the game
			switch cmd {
			case StopGame:
				fmt.Printf("operator quitting the game at address %s\n", gop.contractAddress)
				return
			default:
				fmt.Printf("Unknown command")
			}

		// The default behaviour is to continue operating the game
		default:
			// Proceed with the normal game logic
			gop.operate()
		}
		time.Sleep(1 * time.Second)
	}
}

// This function is the heart of operator. Does the following:
// 1) Figure out in which state the game is
// 2) Force a state transition or payout
func (gop *GameOperator) operate() {
	fmt.Printf("wow look at me I am operating hard %s\n", gop.contractAddress)
}

// Stop stops the operator from operating the game. The game ideally should also
// be paused.
func (gop *GameOperator) Stop() error {
	if gop.playing == false {
		return errors.New("game already stopped")
	}
	gop.controlChannel <- StopGame
	gop.playing = false
	return nil
}
