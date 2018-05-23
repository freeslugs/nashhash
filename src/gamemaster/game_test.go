package gm

import (
	"fmt"
	"testing"
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
	assertEqual(t, g.State(), REVEAL_STATE, "")
	assertEqual(t, g.CurrCommits(), uint(1), "")

	// Test sending another commit into a REVEAL_STATE
	e = g.SendCommitSafe()
	assertNotEqual(t, e, nil, "Should have delivered an error")
	assertEqual(t, g.State(), REVEAL_STATE, "")
	assertEqual(t, g.CurrCommits(), uint(1), "")
}

func TestPlay(t *testing.T) {
	var g Game
	g.Init("0x1", 10)

}
