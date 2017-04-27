package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

var stop bool

func listen(bindAddress string, port string) (ln net.Listener, err error) {
	var buffer bytes.Buffer

	buffer.WriteString(bindAddress)
	buffer.WriteString(":")
	buffer.WriteString(port)

	ln, err = net.Listen("tcp", buffer.String())
	return
}

// Start doc
func Start(bindAddress string, port string, socket string) {
	log.Printf("Starting server using socket '%s' at %s:%s\n", socket, bindAddress, port)

	ln, err := listen(bindAddress, port)

	if err != nil {
		log.Fatalf("Error - %s\n", err.Error())
		return
	}

	for !stop {
		fmt.Print('.')
	}

	ln.Close()
	log.Println("Server is stopped")
}

// Stop doc
func Stop() {
	log.Println("Stopping server")
	stop = true
}

func handleSocket() {

}
