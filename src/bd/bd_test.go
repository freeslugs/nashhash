package bd

import (
	"log"
	"net/rpc"
	"strconv"
	"testing"
)

const (
	GameContract = "0x8bcd426baa7a24e590a9cc65de7a273257163c35"
	OwnerHexKey  = "76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3"
	BotAllowance = 0.2
	RPCPort      = 57543
	RPCAddr      = "127.0.0.1"
)

func TestBotDispatcherRPC(t *testing.T) {
	var bd BotDispatcher
	bd.Init(RPCAddr, RPCPort, OwnerHexKey)
	log.Println("init returned")
	defer bd.Kill()

	bdAddr := RPCAddr + ":" + strconv.Itoa(RPCPort)
	log.Println(bdAddr)

	c, err := rpc.Dial("tcp", bdAddr)
	if err != nil {
		log.Fatal(err)
	}

	args := DispatchArgs{ContractAddress: GameContract, BotAllowance: BotAllowance, Number: 3}
	reply := &DispatchReply{}

	e := call(c, "BotDispatcher.Dispatch", args, reply)
	if e != nil {
		log.Fatal(err)
	}

}
