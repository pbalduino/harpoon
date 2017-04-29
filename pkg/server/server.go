package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

func dockerComm(socket string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	log.Println("Waiting for events")

	eventCh, _ := cli.Events(context.Background(), types.EventsOptions{})

	for event := range eventCh {
		log.Println(event)
		log.Println(event.Status)
		byt := []byte(event.Status)

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}

		// log.Println(dat[])
	}
}

func processMessage(msg string) {
	log.Print("Message:", msg)
}

// Start doc
func Start(bindAddress string, port string, socket string) {
	log.Printf("Starting server using socket '%s' at %s:%s\n", socket, bindAddress, port)

	go clientComm(bindAddress, port)
	go dockerComm(socket)

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
