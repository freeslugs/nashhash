package gm

import (
	"fmt"
	"net/rpc"
	"strconv"
	"testing"
)

func TestGameMasterRPC(t *testing.T) {
	var gm GameMaster

	gm.Init()
	defer gm.Kill()

	var (
		addr     = ":" + strconv.Itoa(11112)
		request  = &ExecuteCallArgs{Message: "test test test"}
		response = new(ExecuteCallReply)
	)

	// Establish the connection to the adddress of the
	// RPC server
	c, _ := rpc.Dial("tcp", addr)
	defer c.Close()

	fmt.Println("client succesfull dialed the number...")

	// Perform a procedure call (core.HandlerName == Handler.Execute)
	// with the Request as specified and a pointer to a response
	// to have our response back.
	_ = c.Call("GameMaster.Execute", request, response)
	fmt.Println(response.Response)

}
