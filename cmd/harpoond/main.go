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
		bindAddress, port, socket string
	)

	flag.StringVar(&socket, "socket", "unix:///var/run/docker.sock", "Socket location")
	flag.StringVar(&port, "port", "10608", "TCP port for client connection")
	flag.StringVar(&bindAddress, "bind", "0.0.0.0", "Bind address for client connection")

	flag.Parse()

	envVars(&socket, &port, &bindAddress)

	serve(socket, fmt.Sprint(port), bindAddress)
}

func envVars(socket *string, port *string, bindAddress *string) {
	hs := os.Getenv("HARPOON_SOCKET")
	if hs != "" {
		*socket = hs
	}

	hp := os.Getenv("HARPOON_PORT")
	if hp != "" {
		*port = hp
	}

	hb := os.Getenv("HARPOON_BIND")
	if hb != "" {
		*bindAddress = hb
	}
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
