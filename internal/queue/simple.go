package queue

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type msgWithSync struct {
	item msg.Msg
	wg   *sync.WaitGroup
}

type SimpleQueue struct {
	items       []msgWithSync
	name        string
	subscribers []string
	handler     msg.Callback
}

func (q SimpleQueue) FindById(id uuid.UUID) (msg.Msg, *sync.WaitGroup) {
	for _, element := range q.items {
		if element.item.GetId() == id {
			return element.item, element.wg
		}
	}
	return nil, nil
}

//Add item to end of queue
func (q *SimpleQueue) AddMsg(m msg.Msg) {
	parentId := m.GetParentId()
	var wgParent *sync.WaitGroup = nil
	if parentId != uuid.Nil {
		_, wgParent = q.FindById(parentId)
	}
	wgSelf := sync.WaitGroup{}
	wgSelf.Add(1)
	newItem := msgWithSync{
		item: m,
		wg:   &wgSelf,
	}
	q.items = append(q.items, newItem)
	log.Print("Added item to  queue :", q.name)
	go m.Process(wgParent, &wgSelf, q.handler, q.subscribers)
}

func (q *SimpleQueue) AddCallback(callback msg.Callback, endpoint string) {
	q.subscribers = append(q.subscribers, endpoint)
	q.handler = callback
}

func NewSimpleQueue(name string) Queue {
	return &SimpleQueue{
		items: make([]msgWithSync, 0),
		name:  name,
	}
}
