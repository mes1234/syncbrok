package queue

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items []msg.Msg
	name  string
}

func (q SimpleQueue) GetItems() []msg.Msg {
	return q.items
}

//Add item to end of queue
func (q SimpleQueue) AddMsg(m msg.Msg) uuid.UUID {
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)
	go m.Process()
	return m.GetId()
}

func NewSimpleQueue(name string) Queue {
	return SimpleQueue{
		items: make([]msg.Msg, 0),
		name:  name,
	}
}
