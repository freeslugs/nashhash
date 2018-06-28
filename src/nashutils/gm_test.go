package nashutils

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GameContract            = "0x8bcd426baa7a24e590a9cc65de7a273257163c35"
	ZeroStageLengthContract = "0xea6c500ce063436e444a4dc7acdd77720bbfe70e"
	OwnerHexKey             = "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	OwnerAddr               = "0x537CA571AEe8116575E8d7a79740c70f685EC856"
	StakeSize               = 10000000000000000 // 0.01 ETH

	DispatcherHexKey = "4a6fd76e5dd2980266a241e23911a6b5870671d3475ed28a04eeadedc7082b6a"
	DispatcherAddr   = "0xa8dAAD283Ca538a3F27371a6f944a4Fa66025957"

	RPCPort = 57543
	RPCAddr = "127.0.0.1"
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

func TestDispatch(t *testing.T) {

	var bd BotDispatcher
	e := bd.Init(RPCAddr, BDPort, DispatcherHexKey, true)
	if e != nil {
		log.Fatal(e)
	}

	var gm GM
	e = gm.Init(RPCAddr, GMPort, OwnerHexKey, false, &bd)
	if e != nil {
		log.Fatal(e)
	}

	//defer gm.Kill()
	gmAddr := RPCAddr + ":" + strconv.Itoa(GMPort)

	var clerk Clerk
	e = clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}

	//defer clerk.Kill()

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

	time.Sleep(50 * time.Second)

	// Connect the game, start the operator
	clerk.ConnectGame(GameContract)
	log.Printf("game connected")

	time.Sleep(5 * time.Second)

	bdAddr := RPCAddr + ":" + strconv.Itoa(BDPort)
	c, err := rpc.Dial("tcp", bdAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Lets ask for a dispatch
	args1 := DispatchArgs{
		ContractAddress: GameContract,
		RequiredBalance: 0.01,
		Number:          2}
	reply := &DispatchReply{}

	e = call(c, "BotDispatcher.Dispatch", args1, reply)
	if e != nil {
		log.Fatalln(e)
	}

	time.Sleep(6 * time.Minute)

	gm.Kill()
	bd.Kill()

	time.Sleep(2 * time.Minute)

}
