package bd

import (
	"log"
	"net/rpc"
)

// DispatchArgs is the argument to the BotDispatcher.Dispatch rpc call
type DispatchArgs struct {
	// ContractAddress is the address of the contract the bot operates on
	ContractAddress string

	// Bot allowance is the balance that the bot has to have to be able
	// to perfoma DoBotStuff
	RequiredBalance float64

	// Number is the number of bots that needs to be dispatched
	Number uint
}

// DispatchReply has the reply
type DispatchReply struct{}

func call(c *rpc.Client, rpcname string,
	args interface{}, reply interface{}) error {

	err := c.Call(rpcname, args, reply)
	if err == nil {
		return nil
	}
	log.Println(err)
	return err
}
