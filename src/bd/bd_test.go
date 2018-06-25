package bd

import (
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	GameContract    = "0x8bcd426baa7a24e590a9cc65de7a273257163c35"
	OwnerHexKey     = "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	OwnerAddr       = "0x537CA571AEe8116575E8d7a79740c70f685EC856"
	RequiredBalance = 0.2
	RPCPort         = 57543
	RPCAddr         = "127.0.0.1"
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

func TestBotDispatcherRPC(t *testing.T) {
	var bd BotDispatcher
	bd.Init(RPCAddr, RPCPort, OwnerHexKey, true)

	// bdAddr := RPCAddr + ":" + strconv.Itoa(RPCPort)

	// _, err := rpc.Dial("tcp", bdAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	time.Sleep(40 * time.Second)

	bd.Kill()

	time.Sleep(30 * time.Second)

}

func TestFindBalance(t *testing.T) {
	var bd BotDispatcher

	bd.queues = make(map[float64]*BotQ)

	bd.queues[1] = nil
	bd.queues[0.1] = nil
	bd.queues[0.5] = nil
	bd.queues[0.8] = nil

	assertEqual(t, bd.findRightBalance(0.99), 1.0, "wrong balance 1")
	assertEqual(t, bd.findRightBalance(0.0001), 0.1, "wrong balance 2")
	assertEqual(t, bd.findRightBalance(100), 0.0, "wrong balance 3")
	assertEqual(t, bd.findRightBalance(0.5), 0.5, "wrong balance 4")
	assertEqual(t, bd.findRightBalance(0.7), 0.8, "wrong balance 5")

}

func TestBotQInit(t *testing.T) {

	var bq BotQ
	botn := uint(10)
	amount := 0.001
	bq.Init(amount, botn)
	defer bq.Kill(common.HexToAddress(OwnerAddr))

	assertEqual(t, len(bq.refill), int(botn), "incoorect initialization")
	assertEqual(t, bq.guaranteedBalance, amount, "incorrect guarantee")

	time.Sleep(3 * time.Second)

}

func TestUtils(t *testing.T) {

	eth := big.NewInt(1000000000000000000)
	pointOne := big.NewInt(100000000000000000)

	ethf := 1.0
	pointOnef := 0.1

	assertEqual(t, eth.Cmp(toWei(1.0)), 0, "incorrect toWei conversion")
	assertEqual(t, pointOne.Cmp(toWei(0.1)), 0, "incorrect toWei conversion")

	assertEqual(t, toEth(eth), ethf, "incorrect toEth conversion")
	assertEqual(t, toEth(pointOne), pointOnef, "incorrect toEth conversion")

	temp := [...]uint{1, 2, 3}

	all, nothing := temp[:3], temp[3:len(temp)]

	log.Println(all)
	log.Println(nothing)

}
