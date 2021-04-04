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
	universe, newMsgCh, newQueuesCh, newSubscribersCh := space.New()
	go universe.Start()

	go frontend.HttpNewMsgController(handleMessage, universe, newMsgCh, newSubscribersCh)()
	go frontend.HttpNewQueueController(newQueuesCh)()
	for {
		time.Sleep(1 * time.Hour)
	}
}
