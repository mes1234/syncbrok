package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items []msg.Msg
}

func (q SimpleQueue) GetItems() []msg.Msg {
	return q.items
}

//Add item to end of queue
func (q SimpleQueue) Add(m msg.Msg) uuid.UUID {
	q.items = append(q.items, m)
	return m.GetId()
}

func NewSimpleQueue() Queue {
	return SimpleQueue{
		items: make([]msg.Msg, 0),
	}
}
