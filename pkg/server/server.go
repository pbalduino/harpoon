package server

import (
	"bytes"
	"log"

	"github.com/pbalduino/harpoon/pkg/client"
	"github.com/pbalduino/harpoon/pkg/docker"
)

var done = make(chan bool)

// Start doc
func Start(bindAddress string, port string, socket string) {
	var buffer bytes.Buffer

	buffer.WriteString(bindAddress)
	buffer.WriteString(":")
	buffer.WriteString(port)

	address := buffer.String()

	go client.Start(address, done)
	go docker.Start(socket)

	for !<-done {
	}

	log.Println("Server is stopped")
}

// Stop doc
func Stop() {
	log.Println("Stopping server")
	done <- true
}

func handleSocket() {

}
