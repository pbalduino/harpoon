package client

import (
	"log"
	"net"
)

// Start does
func Start(address string, done chan bool) {
	socket, err := listen(address)
	var count int

	if err != nil {
		log.Fatalf("Error - %s\n", err.Error())
		return
	}

	defer socket.Close()

	for {
		select {
		case d := <-done:
			if d {
				return
			}
		default:
			conn, err := socket.Accept()
			if err != nil {
				panic(err)
			}

			go handle(conn, count)
			count++
		}
	}
}

func listen(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}

func handle(conn net.Conn, count int) {
	bs := []byte("Sup my friend\n")

	defer conn.Close()

	conn.Write(bs)

	log.Printf("..Closing socket connection #%d\n", count)
}

func process(msg string) {
	log.Print("Message: ", msg)
}
