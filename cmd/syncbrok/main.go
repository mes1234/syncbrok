package main

import (
	"github.com/mes1234/syncbrok/internal/frontend"

	"github.com/mes1234/syncbrok/internal/space"
)

func main() {

	// Begining of life
	universe, newMsgCh, newQueuesCh, newSubscribersCh := space.New()
	go universe.Start()

	frontend.
		HttpStart(
			newMsgCh,
			newSubscribersCh,
			newQueuesCh,
			frontend.HttphandleMessage)

}
