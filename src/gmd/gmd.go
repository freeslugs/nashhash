package main

import (
	"flag"
	"gm"
	"log"
	"net"
	"strconv"
)

func main() {

	// Get the flags
	ipp := flag.String("ip", "127.0.0.1", "the ip address for the game master to listen for connections on")
	portp := flag.String("port", "50000", "the port the game master listens for connections on: 49151 <= port <= 65535 ")
	hexkeyp := flag.String("key", "dead", "secret key of the owner address")
	flag.Parse()

	ip := *ipp
	port := *portp
	portnum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Usage: %v\n", err)
	}
	hexkey := *hexkeyp

	checkInput(ip, port)

	if hexkey == "dead" {
		log.Fatalf("Usage: must provide a key. ./gmd -h for help\n")
	}

	// Init the game master
	var gm gm.GM
	e := gm.Init(ip, portnum, hexkey, false)
	if e != nil {
		log.Fatal(e)
	}
	defer gm.Kill()
	log.Printf("INFO gmd: gamemaster initialized succesfully on %s:%s\n", ip, port)

	// Now we siply sit here indefinately
	select {}

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
