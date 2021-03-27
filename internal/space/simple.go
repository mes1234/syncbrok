package space

import (
	"github.com/mes1234/syncbrok/internal/queue"
)

type SimpleSpace struct {
	queues map[string]queue.Queue
}

func (s SimpleSpace) AddQueue(q queue.Queue, name string) {
	s.queues[name] = q
}

func NewSpace() Space {
	return SimpleSpace{
		queues: make(map[string]queue.Queue),
	}
}
