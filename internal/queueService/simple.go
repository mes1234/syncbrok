package queueService

import (
	"log"
	"time"

	"github.com/mes1234/syncbrok/internal/space"
)

type SimpleQueueService struct {
	handler func()
	space   space.Space
}

func simpleHandler() {
	for {
		time.Sleep(2 * time.Second)
		log.Print("Hello, log file!")
	}
}

func NewSimpleQueueService() SimpleQueueService {
	return SimpleQueueService{
		handler: simpleHandler,
	}
}

func (qs SimpleQueueService) Start(space space.Space) {
	qs.space = space
	go qs.handler()
}
