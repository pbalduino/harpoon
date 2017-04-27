package main

import (
	"fmt"

	"github.com/pbalduino/harpoon/pkg/docker"
)

func main() {
	fmt.Println("I'm the client")
	docker.Hello()
	docker.GoodBye()
}
