package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items []msg.IMsg
}

func (q SimpleQueue) GetItems() []msg.IMsg {
	return q.items
}

//Add item to end of queue
func (q SimpleQueue) Add(m msg.IMsg) uuid.UUID {
	q.items = append(q.items, m)
	return m.GetId()
}

func New() SimpleQueue {
	return SimpleQueue{
		items: make([]msg.IMsg, 0),
	}
}
