package queue

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items    []msg.Msg
	channels []chan bool
	name     string
}

func (q SimpleQueue) GetItems() []msg.Msg {
	return q.items
}

func (q SimpleQueue) FindById(id uuid.UUID) (msg.Msg, chan bool) {
	for index, element := range q.items {
		if element.GetId() == id {
			return element, q.channels[index]
		}
	}
	return nil, nil
}

//Add item to end of queue
func (q *SimpleQueue) AddMsg(m msg.Msg) uuid.UUID {
	parentId := m.GetParentId()
	var parentChan chan bool = nil
	if parentId != uuid.Nil {
		_, parentChan = q.FindById(parentId)
	}
	ch := make(chan bool, 1)
	q.channels = append(q.channels, ch)
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)
	go m.Process(parentChan, ch)
	return m.GetId()
}

func NewSimpleQueue(name string) Queue {
	return &SimpleQueue{
		items:    make([]msg.Msg, 0),
		channels: make([]chan bool, 0),
		name:     name,
	}
}
