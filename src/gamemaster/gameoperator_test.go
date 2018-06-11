package gm

import (
	"testing"
	"time"
)

func TestOperatorInit(t *testing.T) {
	var gop GameOperator
	gm := &GM{debug: true}
	gop.Init("0x1", gm)
	assertEqual(t, gop.ContractAddress(), "0x1", "Wrong contract address")
	assertEqual(t, gop.Playing(), false, "Bad state: should not be playing")
}

func TestPlay(t *testing.T) {
	var gop GameOperator
	gm := &GM{debug: true}

	gop.Init("0x1", gm)

	err := gop.Play()
	if err != nil {
		panic("bad")
	}

	assertEqual(t, gop.Playing(), true, "Bad state: should be playing")
	time.Sleep(2 * time.Second)

	gop.Stop()
	time.Sleep(2 * time.Second)

	assertEqual(t, gop.Playing(), false, "Bad state: should not be playing")
}

func TestRepeatedPlay(t *testing.T) {
	var gop GameOperator
	gm := &GM{debug: true}

	gop.Init("0x1", gm)

	// Call play twice
	err := gop.Play()
	if err != nil {
		panic("bad")
	}

	err = gop.Play()
	assertNotEqual(t, err, nil, "Should not allow repeated play")
	assertEqual(t, gop.Playing(), true, "Bad state: should be playing")

	time.Sleep(2 * time.Second)

	err = gop.Stop()
	assertEqual(t, err, nil, "Should not allow repeated stop")
	assertEqual(t, gop.Playing(), false, "Bad state: should not be playing")

	err = gop.Stop()
	assertEqual(t, gop.Playing(), false, "Bad state: should not be playing")
	assertNotEqual(t, err, nil, "Should not allow repeated stop")

}
