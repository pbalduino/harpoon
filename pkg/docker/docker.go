package docker

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/pbalduino/harpoon/pkg/event"
)

// Start does
func Start(socket string) {
	log.Println("Hello docker")
	parts := strings.Split(socket, "://")

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", parts[1])
			},
		},
	}

	resp, _ := client.Get("http://docker/events")
	reader := bufio.NewReader(resp.Body)

	for {
		line, _ := reader.ReadBytes('\n')

		go event.Process(line)
	}

}
