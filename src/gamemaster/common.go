package gm

import (
	"fmt"
	"net/rpc"
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

func call(srv string, rpcname string,
	args interface{}, reply interface{}) error {
	c, err := rpc.Dial("tcp", srv)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return nil
	}

	fmt.Println(err)
	return err
}

// connectGame. Asks gm at gmAddr to connect game at contract address gameAddr
func connectGame(gmAddr string, gameAddr string) error {
	args := ConnectCallArgs{ContractAddress: gameAddr}
	reply := &ConnectCallReply{}

	e := call(gmAddr, "GM.Connect", args, reply)
	if e != nil {
		return e
	}
	return nil
}

func disconnectGame(gmAddr string, gameAddr string) error {
	args := DisconnectCallArgs{ContractAddress: gameAddr}
	reply := &DisconnectCallReply{}

	e := call(gmAddr, "GM.Disconnect", args, reply)
	if e != nil {
		return e
	}
	return nil
}
