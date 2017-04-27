package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	socketPtr := flag.String("socket", "/var/run/docker.sock", "Socket location")
	portPtr := flag.Int("port", 10608, "TCP port for client connection")
	bindAddressPtr := flag.String("bind", "0.0.0.0", "Bind address for client connection")

	flag.Parse()

	err := checkIfUnique()
	if err != nil {
		fmt.Println("Server is already running")
	} else {
		fmt.Printf("Starting server using socket '%s' at %s:%d\n", *socketPtr, *bindAddressPtr, *portPtr)
	}
}

func checkIfUnique() (err error) {
	_, err = net.Listen("tcp", ":10608")
	return
}
