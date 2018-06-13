package gm

import (
	"fmt"
	"log"
	"strconv"
	"sync"
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

func TestGMRPC(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()

	var (
		addr     = ":" + strconv.Itoa(11112)
		request  = ExecuteCallArgs{Message: "test test test"}
		response = &ExecuteCallReply{}
	)

	e := call(addr, "GM.Execute", request, response)
	if e != nil {
		panic("rpc failed")
	}

	fmt.Println(response.Response)
}

func TestConnectGame(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")

	//time.Sleep(2 * time.Second)

}

func TestRepeatedConnectGame(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	e := connectGame(gmAddr, "0x1")
	assertEqual(t, e, nil, "connect game failed")

	e = connectGame(gmAddr, "0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	//time.Sleep(2 * time.Second)

}

func TestMultipleConnects(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
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
	var gm GM
	gm.Init("", 11112, "", true)
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
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
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
	var gm GM
	gm.Init("", 11112, "", true)
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

func TestBasicThreaded(t *testing.T) {
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)
	numThreads := 10

	var wg sync.WaitGroup
	wg.Add(numThreads)

	// Connect bunch of games, each connection goes in its own thread
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			e := connectGame(gmAddr, "0x"+strconv.Itoa(i))
			assertEqual(t, e, nil, "connect game failed")
			wg.Done()
		}(i)
	}

	wg.Wait()
	//time.Sleep(2 * time.Second)
	assertEqual(t, len(gm.operatedGames), numThreads, "gm mapping has wrong size")

	// Disconnect all these games in async
	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			e := disconnectGame(gmAddr, "0x"+strconv.Itoa(i))
			assertEqual(t, e, nil, "connect game failed")
			wg.Done()
		}(i)
	}

	wg.Wait()
	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

}

func TestClerkInit(t *testing.T) {
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()

	gmAddr := ":" + strconv.Itoa(11112)
	var clerk Clerk
	e := clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}

}

func TestClerkThreaded(t *testing.T) {
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)
	const numThreads = 10

	var clerks [numThreads]Clerk
	for i := 0; i < numThreads; i++ {
		e := clerks[i].Init(gmAddr)
		if e != nil {
			log.Fatal(e)
		}
	}

	var wg sync.WaitGroup
	wg.Add(numThreads)

	// Connect bunch of games, each connection goes in its own thread
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			e := clerks[i].ConnectGame("0x" + strconv.Itoa(i))
			assertEqual(t, e, nil, "connect game failed")
			wg.Done()
		}(i)
	}

	wg.Wait()
	//time.Sleep(2 * time.Second)
	assertEqual(t, len(gm.operatedGames), numThreads, "gm mapping has wrong size")

	// Disconnect all these games in async
	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			e := clerks[i].DisconnectGame("0x" + strconv.Itoa(i))
			assertEqual(t, e, nil, "connect game failed")
			wg.Done()
		}(i)
	}

	wg.Wait()
	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

}

func TestEthereum(t *testing.T) {
	var gm GM
	hexkey := "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	gm.Init("", 11112, hexkey, false)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)
	const numThreads = 10

	clerk := &Clerk{GMAddr: gmAddr}
	clerk.ConnectGame("0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24")

	time.Sleep(5 * time.Second)

}
