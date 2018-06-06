package gm

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"testing"
)

func TestGameMasterRPC(t *testing.T) {
	var gm GameMaster
	gm.Init()

	// Register our baby with net/rpc
	rpcs := rpc.NewServer()
	rpcs.Register(&gm)

	// Create a TCP listener that will listen on `Port`
	l, e := net.Listen("tcp", ":"+strconv.Itoa(11111))
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	gm.l = l

	// Close the listener whenever we stop
	defer l.Close()

	// go gm.l.Accept()
	// time.Sleep(2 * time.Second)
	// fmt.Println("gm is waiting...")

	go func() {
		for gm.dead == false {
			conn, err := gm.l.Accept()
			if err == nil && gm.dead == false {
				go rpcs.ServeConn(conn)
			} else if err == nil {
				conn.Close()
			}
			if err != nil && gm.dead == false {
				fmt.Printf("GameMaster accept: %v\n", err.Error())
				gm.Kill()
			}
		}
	}()
	fmt.Println("gm is waiting...")

	var (
		addr     = "127.0.0.1:" + strconv.Itoa(11111)
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
