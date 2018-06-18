package gm

import (
	"log"
	"net/rpc"
)

// IMPORTANT CONFIGS
const (
	// GameOperator Stuff
	// DisconnectOperator command to stop a game
	DisconnectOperator = 0
	// EthClientPath is the path to the geth.ipc
	EthClientPath = "/Users/me/Library/Ethereum/rinkeby/geth.ipc"

	// BotDispatcher stuff
	NetworkID = 4
)

type ExecuteCallArgs struct {
	Message string
}

type ExecuteCallReply struct {
	Response string
}

// ConnectCallArgs is a struct of arguments to the GM.Connect RPC
type ConnectCallArgs struct {
	ContractAddress string
}

// ConnectCallReply is the reply from GM.Connect RPC
type ConnectCallReply struct {
	Reply string
}

// DisconnectCallArgs is a struct of arguments to the GM.Connect RPC
type DisconnectCallArgs struct {
	ContractAddress string
}

// DisconnectCallReply is the reply from GM.Disconnect RPC
type DisconnectCallReply struct {
	Reply string
}

func call(c *rpc.Client, rpcname string,
	args interface{}, reply interface{}) error {

	err := c.Call(rpcname, args, reply)
	if err == nil {
		return nil
	}

	log.Println(err)
	return err
}
