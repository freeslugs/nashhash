package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gm"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	CmdConnect       = "connect"
	CmdDisconnect    = "disconnect"
	CmdListConnected = "ls"
	CmdQuit          = "quit"
)

func main() {

	// Our flags
	ipp := flag.String("ip", "127.0.0.1", "the ip address for the game master to listen for connections on")
	portp := flag.String("port", "50000", "the port the game master listens for connections on: 49151 <= port <= 65535 ")
	flag.Parse()
	ip := *ipp
	port := *portp

	checkInput(ip, port)

	// Initialize the clerk
	gmAddr := ip + ":" + port
	var clerk gm.Clerk
	e := clerk.Init(gmAddr)
	if e != nil {
		log.Fatal(e)
	}

	interpret(&clerk)

}

func interpret(clerk *gm.Clerk) {
	// Run the io loop
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("gmclerk> ")
		line, e := reader.ReadString('\n')
		if e != nil {
			log.Println("reader failure")
		}
		// convert CRLF to LF
		cmds := strings.Fields(strings.Replace(line, "\n", "", -1))
		if len(cmds) == 0 {
			continue
		}

		switch cmds[0] {
		case CmdConnect:
			addrs, e := getAddresses(cmds)
			if e != nil {
				log.Print(e)
			}
			doCommand(clerk, CmdConnect, addrs)

		case CmdDisconnect:
			addrs, e := getAddresses(cmds)
			if e != nil {
				log.Print(e)
			}
			doCommand(clerk, CmdDisconnect, addrs)

		case CmdListConnected:

		case CmdQuit:
			clerk.Kill()
			log.Print("INFO gmc: done\n")
			return

		default:
			log.Printf("error: uknown command %s\n", cmds[0])
		}

	}

}

// Performs the command
func doCommand(clerk *gm.Clerk, cmd string, addrs []string) {
	switch cmd {
	case CmdConnect:
		for _, addr := range addrs {
			clerk.ConnectGame(addr)
		}
	case CmdDisconnect:
		for _, addr := range addrs {
			clerk.DisconnectGame(addr)
		}

	default:
		panic("something wrong bigtime")
	}
}

// Function get addresses from a file where each line has an address.
func getAddressesFromFile(filename string) ([]string, error) {

	var addrs []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		addr := strings.Replace(line, "\n", "", -1)
		addrs = append(addrs, addr)
	}

	return addrs, nil

}

func getAddresses(cmds []string) ([]string, error) {

	// We need an address
	if len(cmds) < 2 {
		return nil, errors.New("error: command requires a target game address")
	}

	// Are we given a file?
	if cmds[1] == "-f" {
		if len(cmds) != 3 {
			return nil, errors.New("error: need a filename")
		}

		addrs, err := getAddressesFromFile(cmds[2])
		if err != nil {
			log.Print(err)
			return nil, err
		}
		return addrs, nil
	}

	var addrs []string
	addrs = append(addrs, cmds[1])
	return addrs, nil

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
