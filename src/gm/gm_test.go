package gm

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GameContract            = "0x6ff5655c93780ce620dbf35c5cd9c506299d45a9"
	ZeroStageLengthContract = "0xa5bf7aee277b4a035ccb7c7f38d1deebffdb025a"
	OwnerHexKey             = "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	StakeSize               = 100000000000000000 // 0.01 ETH
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

func TestConnectGame(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	// Connect the clerk
	var clerk Clerk
	e := clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer clerk.Kill()

	e = clerk.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	assertEqual(t, len(gm.operatedGames), 1, "must have 1 operating game")

	//time.Sleep(2 * time.Second)

}

func TestRepeatedConnectGame(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var clerk Clerk
	e := clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer clerk.Kill()

	e = clerk.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")

	e = clerk.ConnectGame("0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	//time.Sleep(2 * time.Second)

}

func TestMultipleConnects(t *testing.T) {
	var gm GM

	// Init on localhost and port
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var c Clerk
	e := c.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Kill()

	e = c.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x3")
	assertEqual(t, e, nil, "connect game failed")

	e = c.ConnectGame("0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")
}

func TestDisconnect(t *testing.T) {
	// Init on localhost and port
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var c Clerk
	e := c.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Kill()

	e = c.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	assertEqual(t, len(gm.operatedGames), 1, "gm mapping has wrong size")
	time.Sleep(2 * time.Second)

	e = c.DisconnectGame("0x1")
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

	var c Clerk
	e := c.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Kill()

	// Should fail because no games are connected
	e = c.DisconnectGame("0x1")
	assertNotEqual(t, e, nil, "disconnect game failed")

	e = c.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")

	e = c.DisconnectGame("0x1")
	assertEqual(t, e, nil, "disconnect game failed")

	e = c.DisconnectGame("0x1")
	assertNotEqual(t, e, nil, "repeated connect did not fail")

	//time.Sleep(2 * time.Second)
}

func TestBasic(t *testing.T) {
	// Init on localhost and port
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var c Clerk
	e := c.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Kill()

	e = c.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x3")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect 0x1
	e = c.DisconnectGame("0x1")
	assertEqual(t, e, nil, "disconnect game failed")
	assertEqual(t, len(gm.operatedGames), 2, "gm mapping has wrong size")

	// Reconnect
	e = c.ConnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x1")
	assertNotEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect all
	e = c.DisconnectGame("0x1")
	assertEqual(t, e, nil, "connect game failed")
	e = c.DisconnectGame("0x2")
	assertEqual(t, e, nil, "connect game failed")
	e = c.DisconnectGame("0x3")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

	// Connect 3 more
	e = c.ConnectGame("0x4")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x5")
	assertEqual(t, e, nil, "connect game failed")
	e = c.ConnectGame("0x6")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 3, "gm mapping has wrong size")

	// Disconnect all
	e = c.DisconnectGame("0x4")
	assertEqual(t, e, nil, "connect game failed")
	e = c.DisconnectGame("0x5")
	assertEqual(t, e, nil, "connect game failed")
	e = c.DisconnectGame("0x6")
	assertEqual(t, e, nil, "connect game failed")

	assertEqual(t, len(gm.operatedGames), 0, "gm mapping has wrong size")

}

func TestBasicThreaded(t *testing.T) {
	var gm GM
	gm.Init("", 11112, "", true)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var c Clerk
	e := c.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Kill()

	numThreads := 10

	var wg sync.WaitGroup
	wg.Add(numThreads)

	// Connect bunch of games, each connection goes in its own thread
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			e := c.ConnectGame("0x" + strconv.Itoa(i))
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
			e := c.DisconnectGame("0x" + strconv.Itoa(i))
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
	defer clerk.Kill()

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

func TestKeccak(t *testing.T) {
	guess := []byte("10")
	secret := []byte("3")

	h := crypto.Keccak256(guess, secret)
	fmt.Println(h)
	var hassh [32]byte
	copy(hassh[:], h[:32])
	hash := fmt.Sprintf("0x%x", h)
	res := "0xf3e73e27cf4cda7ebc973aa432f8a54a97200c0745d43e1bc9c2879ffe79cc53"
	assertEqual(t, hash, res, hash)

	fmt.Println(string(guess[:]))
}

func TestEthereumRetrieveState(t *testing.T) {
	var gm GM
	//hexkey := "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(ZeroStageLengthContract), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Game contract: %v", err)
	}

	auth := gm.auth
	assertNotEqual(t, auth, nil, "authenticator should be created")

	stageLength, err := game.GetGameStageLength(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve game stage length: %v", err)
	}
	log.Printf("%v\n", stageLength)

	state, err := game.GetGameStateInfo(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve game state: %v", err)
	}
	log.Printf("%v\n", state)
}

func TestEthereumGetCurrentBlock(t *testing.T) {
	var gm GM
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(header.Number)

}

func TestEthereumBalance(t *testing.T) {
	var gm GM
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	money, err := conn.BalanceAt(context.Background(), gm.auth.From, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(money)
}

func TestEthereumSendETH(t *testing.T) {
	var gm GM
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	// StakeSize = 10000000000000000
	refillAmount := big.NewInt(StakeSize)
	//targetAddress := common.HexToAddress("0x627306090abaB3A6e1400e9345bC60c78a8BEf57")

	// We need to ask the client about currect gas price
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("gas price estimation failed %s\n", err.Error())
	}
	log.Println(gasPrice)

	privk, err := crypto.HexToECDSA(OwnerHexKey)
	assertEqual(t, err, nil, "priv key failed")
	nonce, err := conn.NonceAt(context.Background(), gm.auth.From, nil)
	assertEqual(t, err, nil, "nonce failed")

	// This is the transaction to move money
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress("0x4f77AF2526Ba3a38dA566eADB02Db56075fCe8c1"),
		refillAmount,
		21000, gasPrice, nil)

	log.Println(refillAmount)

	signtx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(4)), privk)
	assertEqual(t, err, nil, "signtx failed")

	errx := conn.SendTransaction(context.Background(), signtx)
	if errx != nil {
		log.Fatalf("tx failed %s\n", errx.Error())
	}

}

func TestEthereumBasic(t *testing.T) {
	var gm GM
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(ZeroStageLengthContract), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Game contract: %v", err)
	}

	auth := gm.auth
	assertNotEqual(t, auth, nil, "authenticator should be created")
	guess := []byte("10")
	secret := []byte("3")

	// Reset the game
	tx, txerr := game.ResetGame(auth)
	assertEqual(t, txerr, nil, "error")

	time.Sleep(1 * time.Second)

	state, err := game.GetGameState(nil)
	if err != nil {
		log.Printf("Failed to retrieve game state: %v", err)
	}
	fmt.Println(state.Int64())

	h := crypto.Keccak256(guess, secret)
	var hash [32]byte
	copy(hash[:], h[:32])

	auth.Value = big.NewInt(StakeSize)
	tx, txerr = game.Commit(auth, hash)
	if txerr != nil {
		log.Fatal(txerr.Error())
	} else {
		log.Printf("commit succesful 0x%x\n", tx.Hash())
	}
	time.Sleep(30 * time.Second)
	auth.Value = nil

	tx, txerr = game.ForceToRevealState(auth)
	if txerr != nil {
		log.Fatalf("Failed to force game into reveal: %v", txerr)
	} else {
		log.Printf("ForceToReveal succesful 0x%x\n", tx.Hash())
	}

	time.Sleep(30 * time.Second)
	state, err = game.GetGameState(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve game state: %v", err)
	}
	fmt.Println(state.Int64())

	time.Sleep(30 * time.Second)
	tx, txerr = game.Reveal(auth, "10", "3")
	if txerr != nil {
		log.Fatal("failed to reveal: %s\n", txerr.Error())
	} else {
		log.Printf("reveal succesful 0x%x\n", tx.Hash())
	}

	time.Sleep(30 * time.Second)
	tx, txerr = game.ForceToPayoutState(auth)
	if txerr != nil {
		log.Fatalf("failed force payout %s\n", txerr.Error())
	} else {
		log.Printf("succesful ForceToPayout 0x%x\n", tx.Hash())
	}

	time.Sleep(30 * time.Second)
	tx, txerr = game.Payout(auth)
	if txerr != nil {
		log.Fatalf("failed payout: %s\n", txerr.Error())
	} else {
		log.Printf("succesful Payout() 0x%x\n", tx.Hash())
	}

}

func TestEthereumGM(t *testing.T) {

	var gm GM
	gm.Init("", 11112, OwnerHexKey, false)
	defer gm.Kill()
	gmAddr := ":" + strconv.Itoa(11112)

	var clerk Clerk
	clerk.Init(gmAddr)
	defer clerk.Kill()

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	game, err := NewGame(common.HexToAddress(GameContract), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Game contract: %v", err)
	}

	// Reset the game
	_, txerr := game.ResetGame(gm.auth)
	assertEqual(t, txerr, nil, "error")

	time.Sleep(20 * time.Second)

	// Connect the game, start the operator
	clerk.ConnectGame(GameContract)

	// Make a commit
	guess := []byte("10")
	secret := []byte("3")
	h := crypto.Keccak256(guess, secret)
	var hash [32]byte
	copy(hash[:], h[:32])
	auth := gm.auth
	auth.Value = big.NewInt(StakeSize)
	tx, txerr := game.Commit(auth, hash)
	if txerr != nil {
		log.Fatal(txerr.Error())
	} else {
		log.Printf("commit succesful 0x%x\n", tx.Hash())
	}
	auth.Value = nil
	time.Sleep(90 * time.Second)

	// Reveal the commit
	tx, txerr = game.Reveal(auth, "10", "3")
	if txerr != nil {
		log.Fatal(txerr.Error())
	} else {
		log.Printf("reveal succesful 0x%x\n", tx.Hash())
	}

	time.Sleep(120 * time.Second)

}
