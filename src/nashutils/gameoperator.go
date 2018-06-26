package nashutils

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GameOperator operates a game contract
type GameOperator struct {
	contractAddress string
	controlChannel  chan int
	playing         bool

	// Bot Dispatcher to run bots
	//bd *BotDispatcher

	// GM controlling us
	gm *GM
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
func (gop *GameOperator) Init(addr string, gm *GM) {
	gop.contractAddress = addr
	gop.controlChannel = make(chan int)
	gop.playing = false
	gop.gm = gm

	// gop.bd = &BotDispatcher{}
	// err := gop.bd.Init(5, gop.contractAddress, gm.auth, gm.key, &gm.authLock)
	// if err != nil {
	// 	log.Printf("ERROR GameOperator %s: init failed %s\n", gop.contractAddress, err.Error())
	// }

	//time.Sleep(10 * time.Second)

}

// Play the game at the contract address
func (gop *GameOperator) Play() error {
	if gop.playing == true {
		log.Printf("WARNING GameOperator: game %s already operated\n", gop.contractAddress)
		return errors.New("GameOperator: game operated")
	}
	go gop.playGame()
	gop.playing = true
	return nil
}

// Game control logic goes into this function
func (gop *GameOperator) playGame() {
	log.Printf("INFO GameOperator %s: playing the game\n", gop.contractAddress)
	for {

		// We use select here for non-blocking read on the socket
		select {
		case cmd := <-gop.controlChannel:

			// Different commands can be passed here.
			// For example, we can ask the handler to reset the state
			// or log the current state of the game
			switch cmd {
			case DisconnectOperator:
				log.Printf("INFO GameOperator %s: quitting the game\n", gop.contractAddress)
				return
			default:
				log.Printf("WARNING GameOperator %s: unknown command\n", gop.contractAddress)
			}

		// The default behaviour is to continue operating the game
		default:
			// Proceed with the normal game logic
			if gop.gm.debug == true {
				log.Printf("INFO GameOperator %s: operate succesful\n", gop.contractAddress)
			} else {
				gop.operate()
			}
		}
	}
}

// This function is the heart of operator. Does the following:
// 1) Figure out in which state the game is
// 2) Force a state transition or payout

const (
	GameCommitState = 0
	GameRevealState = 1
	GamePayoutState = 2
)

// TODO: Not open a new connection every time
func (gop *GameOperator) operate() error {

	defer time.Sleep(15 * time.Second)

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("ERROR GameOperator %s: failed to connect to the Ethereum client: %v", gop.contractAddress, err)
		return err
	}
	defer conn.Close()
	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(gop.contractAddress), conn)
	if err != nil {
		log.Printf("ERROR GameOperator %s: failed to instantiate a Game contract: %v", gop.contractAddress, err)
		return err
	}

	auth := gop.gm.auth

	state, err := game.GetGameStateInfo(nil)
	if err != nil {
		log.Printf("ERROR GameOperator %s: failed to retrieve game state: %v", gop.contractAddress, err)
		return err
	}

	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("ERROR GameOperator %s: failed to get header %v", gop.contractAddress, err)
		return err
	}

	//Initiate appropriate transitions
	switch s := state.State.Int64(); s {
	case GameCommitState:

		// We deploy bots if we are in commiy and 3/4 of gamelength have passed
		// We need to make sure that CommitStageStart block was set this round
		if state.CommitStageStartBlock.IsInt64() {

			// Get the two deadlines
			botDeadline := state.CommitStageStartBlock.Int64() + ((state.StageLength.Int64() * 1) / 2)
			transitionDeadline := state.CommitStageStartBlock.Int64() + state.StageLength.Int64()

			// We can que bots, initiate a transition or do nothing
			if header.Number.Int64() > botDeadline && header.Number.Int64() <= transitionDeadline {
				log.Printf("INFO GameOperator %s: supposeed to add bots but not doing such thing\n", gop.contractAddress)
				//go gop.bd.Dispatch(3)
			} else if header.Number.Int64() > transitionDeadline {
				tx, txerr := game.ForceToRevealState(auth)
				if txerr != nil {
					log.Printf("ERROR GameOperator %s: failed to force game into reveal: %v", gop.contractAddress, txerr)
				} else {
					log.Printf("INFO GameOperator %s: succesful ForceToReveal 0x%x\n", gop.contractAddress, tx.Hash())
				}
			} else {
				log.Printf("INFO GameOperator %s: nothig to be done yet\n", gop.contractAddress)
			}

		} else {
			log.Printf("INFO GameOperator %s: nothing to be done\n", gop.contractAddress)
			return nil
		}

	case GameRevealState:

		if state.RevealStageStartBlock.IsInt64() {

			transitionDeadline := state.RevealStageStartBlock.Int64() + state.StageLength.Int64()

			if header.Number.Int64() > transitionDeadline {

				tx, txerr := game.ForceToPayoutState(auth)
				if txerr != nil {
					log.Printf("ERROR GameOperator %s: failed to force game into payout: %v", gop.contractAddress, txerr)
				} else {
					log.Printf("INFO GameOperator %s: succesful ForceToPayout 0x%x\n", gop.contractAddress, tx.Hash())
				}
			}

		} else {
			log.Printf("INFO GameOperator %s: not ready for payout\n", gop.contractAddress)
		}

	case GamePayoutState:
		tx, txerr := game.Payout(auth)
		if txerr != nil {
			log.Printf("ERROR GameOperator %s: failed to perform payout: %v", gop.contractAddress, txerr)
		} else {
			log.Printf("INFO GameOperator %s: succesful Payout() 0x%x\n", gop.contractAddress, tx.Hash())
		}

	default:
		log.Printf("WARNING GameOperator %s: unknown operation state", gop.contractAddress)

	}

	return nil

}

// Stop stops the operator from operating the game. The game ideally should also
// be paused.
func (gop *GameOperator) Stop() error {
	if gop.playing == false {
		log.Printf("WARNING GameOperator %s: game already stopped\n", gop.contractAddress)
		return errors.New("game already stopped")
	}
	gop.controlChannel <- DisconnectOperator
	gop.playing = false

	// Kill the dispatcher
	// err := gop.bd.Kill()
	// if err != nil {
	// 	return err
	// }
	return nil
}
