package main

import (
	"gm"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	parseInput()
	gmAddr := os.Args[1] + ":" + os.Args[2]

	clerk := &gm.Clerk{GMAddr: gmAddr}
	log.Printf("INFO gmc: estailishing connection to GM on %s\n", gmAddr)

	e := clerk.ConnectGame("0x3")
	if e != nil {
		log.Println(e)
	}
}

func parseInput() {
	args := os.Args
	if len(args) != 3 {
		log.Fatalf("Usage: ./gmc <gm_ip> <port>\n")
	}

	ip := os.Args[1]

	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		log.Fatalln("Usage: provided ip address is invalid")
	}

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Usage: %v\n", err)
	}

	if port < 49151 || port > 65535 {
		log.Fatalln("Usage: port has to be in 49151 – 65535 range")
	}
}
