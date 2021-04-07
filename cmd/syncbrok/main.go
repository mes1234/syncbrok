package main

import (
	"log"

	"github.com/mes1234/syncbrok/internal/frontend"

	"github.com/mes1234/syncbrok/internal/space"
)

func handleMessage(content []byte, endpoint string) bool {
	log.Print("This message content: ", string(content))
	return true
}

func main() {

	// Begining of life
	universe, newMsgCh, newQueuesCh, newSubscribersCh := space.New()
	go universe.Start()

	frontend.HttpNewMsgController(newMsgCh)
	frontend.HttpNewSubscriberController(handleMessage, newSubscribersCh)
	frontend.HttpNewQueueController(newQueuesCh)
	frontend.HttpStart()

}
