package nashutils

import (
	"fmt"
	"log"
	"math/big"
	"net/rpc"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	//GameContract = "0x8bcd426baa7a24e590a9cc65de7a273257163c35"
	//OwnerHexKey     = "4a6fd76e5dd2980266a241e23911a6b5870671d3475ed28a04eeadedc7082b6a"
	//OwnerAddr       = "0x537CA571AEe8116575E8d7a79740c70f685EC856"
	RequiredBalance = 0.2
)

func TestBotDispatcherInit(t *testing.T) {
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

func TestBotDispatch(t *testing.T) {

	var bd BotDispatcher
	bd.Init(RPCAddr, RPCPort, OwnerHexKey, true)
	time.Sleep(60 * time.Second)

	bdAddr := RPCAddr + ":" + strconv.Itoa(RPCPort)
	c, err := rpc.Dial("tcp", bdAddr)
	if err != nil {
		log.Fatal(err)
	}

	args1 := DispatchArgs{ContractAddress: "0x1", RequiredBalance: 0.01, Number: 1}
	//args2 := DispatchArgs{ContractAddress: "0x1", RequiredBalance: 0.01, Number: 2}
	reply := &DispatchReply{}

	e := call(c, "BotDispatcher.Dispatch", args1, reply)
	if e != nil {
		log.Fatalln(e)
	}

}

func TestUtils(t *testing.T) {

	// test toWei and toEth
	eth := big.NewInt(1000000000000000000)
	pointOne := big.NewInt(100000000000000000)

	ethf := 1.0
	pointOnef := 0.1

	assertEqual(t, eth.Cmp(toWei(1.0)), 0, "incorrect toWei conversion")
	assertEqual(t, pointOne.Cmp(toWei(0.1)), 0, "incorrect toWei conversion")

	assertEqual(t, toEth(eth), ethf, "incorrect toEth conversion")
	assertEqual(t, toEth(pointOne), pointOnef, "incorrect toEth conversion")

	// Test the keccak equivalence to solidity
	guess := []byte("10")
	secret := []byte("3")
	h := crypto.Keccak256(guess, secret)
	fmt.Println(h)
	var hassh [32]byte
	copy(hassh[:], h[:32])
	hash := fmt.Sprintf("0x%x", h)
	res := "0xf3e73e27cf4cda7ebc973aa432f8a54a97200c0745d43e1bc9c2879ffe79cc53"
	assertEqual(t, hash, res, hash)

	//fmt.Println(string(guess[:]))

}
