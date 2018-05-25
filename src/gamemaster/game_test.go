package gm

import (
	"fmt"
	"testing"
	"time"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertState(t *testing.T, g *Game, state uint, commits uint, reveals uint) {
	assertEqual(t, g.State(), state, "Bad state transition")
	assertEqual(t, g.CurrReveals(), reveals, "Wrong number of reveals")
	assertEqual(t, g.CurrCommits(), commits, "Wrong number of commits")
}

func sendCommits(g *Game, num uint) {
	for i := 0; uint(i) < num; i++ {
		g.SendCommitSafe()
	}
}

func sendReveals(g *Game, num uint) {
	for i := 0; uint(i) < num; i++ {
		g.SendRevealSafe()
	}
}

func TestInit(t *testing.T) {

	var g Game
	g.Init("0x1", 10)

	assertEqual(t, g.MaxPlayers(), uint(10), "Wrong max players")
	assertEqual(t, g.Address(), "0x1", "Wrong contract address")
	assertEqual(t, g.State(), uint(0), "")
}

func TestSendCommit(t *testing.T) {
	var g Game
	g.Init("0x1", 1)

	// Send a single commit. Should change state.
	e := g.SendCommitSafe()
	assertEqual(t, e, nil, "")
	assertEqual(t, g.State(), REVEAL_STATE, "Bad state")
	assertEqual(t, g.CurrCommits(), uint(1), "")

	// Test sending another commit into a REVEAL_STATE
	e = g.SendCommitSafe()
	assertNotEqual(t, e, nil, "SendCommit at wrong state has to return error")
	assertEqual(t, g.State(), REVEAL_STATE, "Bad state")
	assertEqual(t, g.CurrCommits(), uint(1), "")
}

func TestSendReveal(t *testing.T) {
	var g Game
	g.Init("0x1", 1)

	// Send the fist commit
	g.SendCommitSafe()

	// Send reveal
	e := g.SendRevealSafe()
	assertEqual(t, e, nil, "")
	assertState(t, &g, PAYOUT_STATE, 1, 1)

	e = g.SendRevealSafe()
	assertNotEqual(t, e, nil, "SendReveal at wrong state has to return error")
	assertState(t, &g, PAYOUT_STATE, 1, 1)
}

func TestPayout(t *testing.T) {
	var g Game
	g.Init("0x1", 2)

	// Send 2 commits and
	g.SendCommitSafe()
	g.SendCommitSafe()

	// Send 2 reveals
	g.SendRevealSafe()
	g.SendRevealSafe()

	// Assert correct state
	assertState(t, &g, PAYOUT_STATE, 2, 2)

	// Call payout
	e := g.Payout()
	assertEqual(t, e, nil, "")
	assertState(t, &g, COMMIT_STATE, 0, 0)
}

func TestForceToRevealState(t *testing.T) {
	nump := uint(10)
	var g Game
	g.Init("0x1", nump)

	// Send the commits
	sendCommits(&g, 7)
	assertState(t, &g, COMMIT_STATE, 7, 0)

	g.ForceToRevealStateSafe()
	assertState(t, &g, REVEAL_STATE, 7, 0)

	sendCommits(&g, 7)
	assertState(t, &g, REVEAL_STATE, 7, 0)

	e := g.ForceToRevealStateSafe()
	assertNotEqual(t, e, nil, "Bad state")
}

func TestForceToPayoutState(t *testing.T) {
	nump := uint(10)
	var g Game
	g.Init("0x1", nump)

	sendCommits(&g, 7)
	assertState(t, &g, COMMIT_STATE, 7, 0)
	g.ForceToRevealStateSafe()
	assertState(t, &g, REVEAL_STATE, 7, 0)

	sendReveals(&g, 5)
	assertState(t, &g, REVEAL_STATE, 7, 5)

	g.ForceToPayoutStateSafe()
	assertState(t, &g, PAYOUT_STATE, 7, 5)

}

func TestBlockTicker(t *testing.T) {
	nump := uint(10)

	var g Game
	g.Init("0x1", nump)
	quit := make(chan bool, 1)

	go g.BlockTicker(uint(10), quit)

	time.Sleep(1 * time.Second)

	// Check that the block number is higher than 80
	blockNum := g.CurrBlockNumber()
	fmt.Printf("block number %d\n", blockNum)
	good := g.CurrBlockNumber() > uint(50)

	assertEqual(t, true, good, "")

	quit <- true
}

func TestPlayBasic(t *testing.T) {

	nump := uint(10)

	var g Game
	g.Init("0x1", nump)
	toGame := make(chan bool, 1)

	// Start the game
	go g.PlayBasic(toGame)

	// We want to play 5 full rounds
	for i := 0; i < 5; i++ {

		// Wait for the game to reach the payout state
		for g.State() != PAYOUT_STATE {
		}

		// Check that we are ready for payout
		assertState(t, &g, PAYOUT_STATE, nump, nump)

		// Call payout
		g.PayoutSafe()

		// Check if payout has reset us correctly
		assertState(t, &g, COMMIT_STATE, 0, 0)

		toGame <- true
	}
	toGame <- false
}

func TestPlayHard(t *testing.T) {

	nump := uint(10)
	sleepTime := 3000

	var g Game
	g.Init("0x1", nump)
	toGame := make(chan bool, 1)

	// Start the game
	go g.PlayHard(toGame, 70)

	for {
		state := g.State()
		switch state {
		case COMMIT_STATE:
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			if g.State() == COMMIT_STATE {
				g.ForceToRevealStateSafe()
			}
		case REVEAL_STATE:
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			if g.State() == REVEAL_STATE {
				g.ForceToPayoutStateSafe()
			}
		case PAYOUT_STATE:

			assertEqual(t, g.State(), PAYOUT_STATE, "")
			e := g.PayoutSafe()
			assertEqual(t, e, nil, "")
			return

		}

	}

}
