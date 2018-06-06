package gm

import (
	"fmt"
	"net/rpc"
	"strconv"
	"testing"
)

func TestGameMasterRPC(t *testing.T) {
	var gm GameMaster

	// Init on localhost and port
	gm.Init("", 11112)
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

	_ = c.Call("GameMaster.Execute", request, response)
	fmt.Println(response.Response)

}
