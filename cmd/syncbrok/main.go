package main

import (
	"log"
	"time"

	"github.com/mes1234/syncbrok/internal/frontend"

	"github.com/mes1234/syncbrok/internal/space"
)

func handleMessage(content []byte) bool {
	log.Print("This message content: ", string(content))
	return true
}

func main() {

	// Begining of life
	universe := space.New()
	handler := frontend.HttpListner(universe, handleMessage)
	handler()
	time.Sleep(10 * time.Second)
}
