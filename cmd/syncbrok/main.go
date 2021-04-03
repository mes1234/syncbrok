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
	universe, newMsgCh, newQueuesCh := space.New()

	handler := frontend.HttpListner(universe, handleMessage, newMsgCh, newQueuesCh)
	handler()
	time.Sleep(10 * time.Second)
}
