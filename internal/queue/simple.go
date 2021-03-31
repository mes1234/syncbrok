package queue

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items      []msg.Msg
	waitGroups []*sync.WaitGroup
	name       string
}

func (q SimpleQueue) GetItems() []msg.Msg {
	return q.items
}

func (q SimpleQueue) FindById(id uuid.UUID) (msg.Msg, *sync.WaitGroup) {
	for index, element := range q.items {
		if element.GetId() == id {
			return element, q.waitGroups[index]
		}
	}
	return nil, nil
}

//Add item to end of queue
func (q *SimpleQueue) AddMsg(m msg.Msg) uuid.UUID {
	parentId := m.GetParentId()
	var wgParent *sync.WaitGroup = nil
	if parentId != uuid.Nil {
		_, wgParent = q.FindById(parentId)
	}
	wgSelf := sync.WaitGroup{}
	wgSelf.Add(1)
	q.waitGroups = append(q.waitGroups, &wgSelf)
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)
	go m.Process(wgParent, &wgSelf)
	return m.GetId()
}

func NewSimpleQueue(name string) Queue {
	return &SimpleQueue{
		items:      make([]msg.Msg, 0),
		waitGroups: make([]*sync.WaitGroup, 0),
		name:       name,
	}
}
