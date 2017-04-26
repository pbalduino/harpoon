package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	containerIDPtr := flag.String("container", "", "a string")
	initPtr := flag.Bool("init", false, "Starts Harpoon as a server. You need to run it before to use Harpoon.")

	flag.Parse()

	if *initPtr {
		err := checkIfUnique()
		if err != nil {
			fmt.Println("Server is already running")
		} else {
			fmt.Println("Starting server")
		}
	} else if *containerIDPtr != "" {
		fmt.Print("Setting client: ")
		fmt.Println(*containerIDPtr)
	} else {
		fmt.Println("You need to inform a container")
		fmt.Println(*containerIDPtr)
	}
}

func checkIfUnique() (err error) {
	_, err = net.Listen("tcp", ":10608")
	return
}
