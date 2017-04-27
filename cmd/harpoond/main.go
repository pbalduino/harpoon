package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pbalduino/harpoon/pkg/server"
)

func main() {
	var (
		socketPtr      = flag.String("socket", "/var/run/docker.sock", "Socket location")
		portPtr        = flag.Uint("port", 10608, "TCP port for client connection")
		bindAddressPtr = flag.String("bind", "0.0.0.0", "Bind address for client connection")
	)

	flag.Parse()

	serve(*socketPtr, fmt.Sprint(*portPtr), *bindAddressPtr)
}

func serve(socket string, port string, bindAddress string) {
	registerInterrupt()

	server.Start(bindAddress, port, socket)
}

func registerInterrupt() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			log.Println("Received SIGINT")
			server.Stop()
		}
	}()
}
