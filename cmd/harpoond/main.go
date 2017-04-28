package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pbalduino/harpoon/pkg/server"
)

func main() {
	var (
		socketPtr      = flag.String("socket", "unix:///var/run/docker.sock", "Socket location")
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
	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-c
		server.Stop()
	}()
}
