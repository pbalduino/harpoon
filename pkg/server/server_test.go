package server_test

import (
	"testing"
	"time"

	"github.com/pbalduino/harpoon/pkg/server"
)

func TestStartStopServer(t *testing.T) {
	go server.Start("0.0.0.0", "10608", "unix:///var/run/docker.sock")

	time.Sleep(1 * time.Second)

	server.Stop()
}
