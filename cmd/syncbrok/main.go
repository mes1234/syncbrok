package main

import (
	"sync"

	"github.com/mes1234/syncbrok/internal/frontend"

	"github.com/mes1234/syncbrok/internal/space"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Begining of life
	universe, newMsgCh, newQueuesCh, newSubscribersCh := space.New()
	go universe.Start(&wg)
	go frontend.HttpStart(&wg, newMsgCh, newSubscribersCh, newQueuesCh)
	wg.Wait()
}
