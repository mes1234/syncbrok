package space

import (
	"log"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

type SimpleSpace struct {
	queues map[string]queue.Queue
}

func (s SimpleSpace) AddQueue(name string) {
	log.Print("Added new queue ", name)
	s.queues[name] = queue.NewSimpleQueue(name)
}

func (s SimpleSpace) Publish(queueName string, m msg.Msg) {
	s.queues[queueName].AddMsg(m)
}

func New() Space {
	return SimpleSpace{
		queues: make(map[string]queue.Queue),
	}
}
