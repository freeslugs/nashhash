package gm

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestGameMasterRPC(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
	defer gm.Kill()

	var (
		addr     = ":" + strconv.Itoa(11112)
		request  = ExecuteCallArgs{Message: "test test test"}
		response = &ExecuteCallReply{}
	)

	e := call(addr, "GameMaster.Execute", request, response)
	if e != nil {
		panic("rpc failed")
	}

	fmt.Println(response.Response)
}

func TestConnectGame(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")

	//time.Sleep(2 * time.Second)

}

func TestRepeatedConnectGame(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")

	e = connectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	//time.Sleep(2 * time.Second)

}

func TestMultipleConnects(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x3")
	assertEqual(t, e, nil, "connect game failed")

	e = connectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")
}

func TestDisconnect(t *testing.T) {
	// Init on localhost and port
	var gm GameMaster
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")
	assertEqual(t, len(gm.operatedGames), 1, "gm mapping has wrong size")
	time.Sleep(2 * time.Second)

	e = disconnectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "disconnect game failed")

	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

	time.Sleep(2 * time.Second)
}

func TestRepeatedDisconnectGame(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	// Should fail because no games are connected
	e := disconnectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "disconnect game failed")

	e = connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")

	e = disconnectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "disconnect game failed")

	e = disconnectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	//time.Sleep(2 * time.Second)
}

func TestBasic(t *testing.T) {
	// Init on localhost and port
	var gm GameMaster
	gm.Init("", 11112)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x3")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect 0x1
	e = disconnectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "disconnect game failed")
	assertEqual(t, len(gm.operatedGames), 2, "gm mapping has wrong size")

	// Reconnect
	e = connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect all
	e = disconnectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = disconnectGame(gmAddr, "0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = disconnectGame(gmAddr, "0x3")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

	// Connect 3 more
	e = connectGame(gmAddr, "0x4")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x5")
	assertEqual(t, e, nil, "connect game failed")
	e = connectGame(gmAddr, "0x6")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect all
	e = disconnectGame(gmAddr, "0x4")
	assertEqual(t, e, nil, "connect game failed")
	e = disconnectGame(gmAddr, "0x5")
	assertEqual(t, e, nil, "connect game failed")
	e = disconnectGame(gmAddr, "0x6")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

}
