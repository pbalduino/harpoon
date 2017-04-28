package server

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"time"
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

func clientComm(bindAddress string, port string) {
	client, err := listen(bindAddress, port)

	if err != nil {
		log.Fatalf("Error - %s\n", err.Error())
		return
	}

	for {
		var buffer bytes.Buffer

		buffer.WriteString("Sup my friend\n")

		conn, err := client.Accept()

		if err != nil {
			time.Sleep(1000)
			log.Fatal(err)
		}

		msg, _ := bufio.NewReader(conn).ReadString('\n')

		go processMessage(msg)

		conn.Close()
	}
}

func processMessage(msg string) {
	log.Print("Message:", msg)
}

// Start doc
func Start(bindAddress string, port string, socket string) {
	log.Printf("Starting server using socket '%s' at %s:%s\n", socket, bindAddress, port)

	go clientComm(bindAddress, port)

	//	dockerd, err := listen(socket)

	for !stop {
	}

	log.Println("Server is stopped")
}

// Stop doc
func Stop() {
	log.Println("Stopping server")
	stop = true
}

func handleSocket() {

}
