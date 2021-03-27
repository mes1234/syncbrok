package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	Items []msg.IMsg
}

//Add item to end of queue
func (q SimpleQueue) Add(m msg.IMsg) uuid.UUID {
	q.Items = append(q.Items, m)
	return m.GetId()
}

func New() SimpleQueue {
	return SimpleQueue{
		Items: make([]msg.IMsg, 0),
	}
}
