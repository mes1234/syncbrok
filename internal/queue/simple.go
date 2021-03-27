package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	Items []msg.Msg
}

//Add item to end of queue
func (q SimpleQueue) Add(m msg.Msg) uuid.UUID {
	q.Items = append(q.Items, m)
	return m.Id
}

func New() SimpleQueue {
	return SimpleQueue{
		Items: make([]msg.Msg, 0),
	}
}
