package main

import (
	"flag"
	"gm"
	"log"
	"net"
	"strconv"
)

func main() {

	// Our flags
	ipp := flag.String("ip", "127.0.0.1", "the ip address for the game master to listen for connections on")
	portp := flag.String("port", "50000", "the port the game master listens for connections on: 49151 <= port <= 65535 ")
	flag.Parse()
	ip := *ipp
	port := *portp

	checkInput(ip, port)

	gmAddr := ip + ":" + port
	var clerk gm.Clerk
	e := clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}
	defer clerk.Kill()

	clerk.ConnectGame("0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24")

}

func checkInput(ip string, port string) {

	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		log.Fatalln("Usage: provided ip address is invalid")
	}

	portnum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Usage: %v\n", err)
	}

	if portnum < 49151 || portnum > 65535 {
		log.Fatalln("Usage: port has to be in 49151 – 65535 range")
	}
}
