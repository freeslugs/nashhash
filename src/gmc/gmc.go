package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Usage: ./gmc <gm_ip>:<port>\n")
	}

	fmt.Println("!oG ,olleH")
}
