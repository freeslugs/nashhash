package gm

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX_TIME_PERIOD = 5
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
}

// Init initializes the game to a contract address and the max players
func (g *Game) Init(addr string, maxp uint) {
	g.address = addr
	g.maxPlayers = maxp
}

// Address returns the game address
func (g *Game) Address() string {
	return g.address
}

// State returns the state of the game
func (g *Game) State() uint {
	return g.state
}

// MaxPlayers returns the max players of the game
func (g *Game) MaxPlayers() uint {
	return g.maxPlayers
}

func (g *Game) CurrCommits() uint {
	return g.currCommits
}

// Play emulates the Game.sol changing of state. Random increment of commits/reveals
// and appropriate state transitions are used to achive this goal.
func (g *Game) Play() {

	rand.Seed(time.Now().UnixNano())

	for {

		// Sleep a random number of seconds
		sleepTime := rand.Intn(MAX_TIME_PERIOD)
		time.Sleep(time.Duration(sleepTime) * time.Second)

		g.gameLock.Lock()

		switch g.state {
		case COMMIT_STATE:
			if g.currCommits < g.maxPlayers {
				g.currCommits++
				if g.currCommits == g.maxPlayers {
					g.state = REVEAL_STATE
				}
			}
		case REVEAL_STATE:
			if g.currReveals < g.maxPlayers {
				g.currReveals++
				if g.currReveals == g.maxPlayers {
					g.state = PAYOUT_STATE
				}
			}
		case PAYOUT_STATE:

		default:
			fmt.Println("Error: Bad game state")
		}

		g.gameLock.Unlock()
	}
}

// SendCommit imitates sending of a commit to the game.
// WARNING: Assumes that the game lock is held
func (g *Game) SendCommit() error {
	if g.state != COMMIT_STATE {
		return errors.New("Bad state: cannot commit state")
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
