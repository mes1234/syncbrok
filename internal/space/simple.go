package space

import (
	"log"

	"github.com/mes1234/syncbrok/internal/queue"
)

type SimpleSpace struct {
	queues map[string]queue.Queue
}

func (s SimpleSpace) AddQueue(q queue.Queue, name string) {
	log.Print("Added new queue ", name)
	s.queues[name] = q
}

func New() Space {
	return SimpleSpace{
		queues: make(map[string]queue.Queue),
	}
}
