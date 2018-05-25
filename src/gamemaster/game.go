package gm

/*
This file is needed to emulate the behaviour of a blockchain game. Main purpose
of such emulation is to be able to quickly develop the rest of the game master.
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX_TIME_PERIOD = 100
const COMMIT_STATE = uint(0)
const REVEAL_STATE = uint(1)
const PAYOUT_STATE = uint(2)

type Game struct {
	address         string
	state           uint
	roundStartBlock uint
	round           uint

	currCommits uint
	currReveals uint
	maxPlayers  uint

	gameLock sync.Mutex

	// For now the block number is part of the game.
	currBlockNumber uint
	commitPeriod    uint
	revealPeriod    uint
	blockLock       sync.Mutex
}

// BlockTicker is run in a separate goroutine. Pass value of true into the quit channel
// to stop the ticker. Funtion increments the length of the blockchain by one ever millisPerBlock
func (g *Game) BlockTicker(millisPerBlock uint, quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Duration(millisPerBlock) * time.Millisecond)
			g.IncrementBlockNumber()
		}
	}

}

// CurrBlockNumber returns the current number of the head of the blockchain
func (g *Game) CurrBlockNumber() uint {
	g.blockLock.Lock()
	defer g.blockLock.Unlock()
	return g.currBlockNumber
}

// IncrementBlockNumber increments the block number. Does so safely by holding the blockLock mutex
func (g *Game) IncrementBlockNumber() {
	g.blockLock.Lock()
	defer g.blockLock.Unlock()
	g.currBlockNumber++
}

// Init initializes the game to a contract address and the max players
func (g *Game) Init(addr string, maxp uint) {
	g.address = addr
	g.maxPlayers = maxp
	g.revealPeriod = 6
	g.commitPeriod = 6
}

// Address returns the game address
func (g *Game) Address() string {
	g.gameLock.Lock()
	defer g.gameLock.Unlock()
	return g.address
}

// State returns the state of the game
func (g *Game) State() uint {
	g.gameLock.Lock()
	defer g.gameLock.Unlock()
	return g.state
}

// MaxPlayers returns the max players of the game
func (g *Game) MaxPlayers() uint {
	g.gameLock.Lock()
	defer g.gameLock.Unlock()
	return g.maxPlayers
}

// CurrCommits returns the current number of commits
func (g *Game) CurrCommits() uint {
	g.gameLock.Lock()
	defer g.gameLock.Unlock()
	return g.currCommits
}

// CurrReveals returns the current number of reveals
func (g *Game) CurrReveals() uint {
	g.gameLock.Lock()
	defer g.gameLock.Unlock()
	return g.currReveals
}

// PlayBasic emulates the Game.sol changing of state. Random increment of commits/reveals
// and appropriate state transitions are used to achive this goal.
func (g *Game) PlayBasic(fromHandler <-chan bool) {

	rand.Seed(time.Now().UnixNano())
	for {

		// Sleep a random number of seconds
		sleepTime := rand.Intn(MAX_TIME_PERIOD)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

		g.gameLock.Lock()

		switch g.state {
		case COMMIT_STATE:
			g.SendCommit()
		case REVEAL_STATE:
			g.SendReveal()
		case PAYOUT_STATE:
			g.gameLock.Unlock()
			stayAlive := <-fromHandler
			if !stayAlive {
				fmt.Println("die")
				return
			}
			continue

			// Do nothing, wait for payout to be called
		default:
			fmt.Println("Error: Bad game state")
			g.gameLock.Unlock()
			return
		}

		g.gameLock.Unlock()
	}
}

// PlayHard emulates the state transition of a game
func (g *Game) PlayHard(fromHandler <-chan bool, chanceIgnore int) {
	rand.Seed(time.Now().UnixNano())

	// Total possible commits
	totalCommits := g.maxPlayers
	totalReveals := g.maxPlayers

	for {

		// Sleep a random number of seconds
		sleepTime := rand.Intn(MAX_TIME_PERIOD)
		ignore := (rand.Intn(100) > chanceIgnore)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

		switch g.State() {
		case COMMIT_STATE:
			if totalCommits == 0 {
				continue
			}

			totalCommits--
			if !ignore {
				g.SendCommitSafe()

			}

		case REVEAL_STATE:
			if totalReveals == 0 {
				continue
			}

			totalReveals--
			if !ignore {
				g.SendRevealSafe()
			}

		case PAYOUT_STATE:
			return

			// Do nothing, wait for payout to be called
		default:
			fmt.Println("Error: Bad game state")
			return
		}
	}

}

// SendCommit imitates sending of a commit to the game.
// WARNING: Assumes that the game lock is held
func (g *Game) SendCommit() error {
	if g.state != COMMIT_STATE {
		return errors.New("Bad state: cannot commit in this state")
	}

	if g.currCommits < g.maxPlayers {
		g.currCommits++
		if g.currCommits == g.maxPlayers {
			g.state = REVEAL_STATE
		}
		return nil
	}
	return errors.New("WRONG BEHAVIOUR: Inconsistent State")
}

// SendCommitSafe sends a commit but also lock the gamelock
// WARNING: Assumes the gamelock is not held
func (g *Game) SendCommitSafe() error {
	g.gameLock.Lock()
	e := g.SendCommit()
	g.gameLock.Unlock()
	return e
}

// SendReveal imitates sending of a reveal to the game.
// WARNING: Assumes that the game lock is held
func (g *Game) SendReveal() error {
	if g.state != REVEAL_STATE {
		return errors.New("Bad state: cannot reveal in this state")
	}

	if g.currReveals < g.currCommits {
		g.currReveals++
		if g.currReveals == g.maxPlayers {
			g.state = PAYOUT_STATE
		}
		return nil
	}
	return errors.New("WRONG BEHAVIOUR: Inconsistent State")
}

// SendRevealSafe sends a reveal but also locks the gamelock
// WARNING: Assumes the gamelock is not held
func (g *Game) SendRevealSafe() error {
	g.gameLock.Lock()
	e := g.SendReveal()
	g.gameLock.Unlock()
	return e
}

// Payout triggers the payout routine in the contract. In emulation, it resets the state to COMMIT_STATE
// WARNING: Assumes that the game lock is held
func (g *Game) Payout() error {
	if g.state != PAYOUT_STATE {
		return errors.New("Bad state: cannot payout in this state")
	}
	g.state = COMMIT_STATE
	g.currCommits = 0
	g.currReveals = 0
	return nil
}

// PayoutSafe like payout but also locks the gamelock
// WARNING: Assume the lock is NOT held
func (g *Game) PayoutSafe() error {
	g.gameLock.Lock()
	e := g.Payout()
	g.gameLock.Unlock()
	return e
}

// ForceToRevealState forces the game into reveal state
// WARNING: Assumes the game lock is held
func (g *Game) ForceToRevealState() error {
	if g.state != COMMIT_STATE {
		return errors.New("Bad state: cannot force to reveal in this state")
	}
	g.state = REVEAL_STATE
	return nil
}

// ForceToRevealStateSafe same but holds the lock
// WARNING: Assumes the game lock is NOT held
func (g *Game) ForceToRevealStateSafe() error {
	g.gameLock.Lock()
	e := g.ForceToRevealState()
	g.gameLock.Unlock()
	return e
}

// ForceToPayoutState forces the game into PAYOUT_STATE
// WARNING: Assumes the game lock is held
func (g *Game) ForceToPayoutState() error {
	if g.state != REVEAL_STATE {
		return errors.New("Bad state: cannot force to reveal in this state")
	}
	g.state = PAYOUT_STATE
	return nil
}

// ForceToPayoutStateSafe same but holds the lock
// WARNING: Assumes the game lock is NOT held
func (g *Game) ForceToPayoutStateSafe() error {
	g.gameLock.Lock()
	e := g.ForceToPayoutState()
	g.gameLock.Unlock()
	return e
}
