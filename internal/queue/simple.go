package queue

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/storage"
)

type SimpleQueue struct {
	items       []msg.Msg
	name        string
	subscribers []string
	handler     msg.Callback
	storage     chan<- msg.Msg
	storageAck  chan<- uuid.UUID
	storeReader storage.FileReader
}

func (q SimpleQueue) FindById(id uuid.UUID) msg.Msg {
	for _, element := range q.items {
		if element.GetId() == id {
			return element
		}
	}
	return nil
}

//Add item to end of queue
func (q *SimpleQueue) AddMsg(m msg.Msg) {

	q.storage <- m
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)

	parentId := m.GetParentId()
	var parent msg.Msg = nil
	if parentId != uuid.Nil {
		parent = q.FindById(parentId)
		go m.Process(parent.GetWaiter(), q.handler, q.subscribers, q.storageAck)
	} else {
		go m.Process(nil, q.handler, q.subscribers, q.storageAck)
	}

}

func (q *SimpleQueue) AddCallback(callback msg.Callback, endpoint string) {
	q.subscribers = append(q.subscribers, endpoint)
	q.handler = callback
}

func NewSimpleQueue(name string, storage chan<- msg.Msg, ackMessageCh chan<- uuid.UUID, storeReader storage.FileReader) Queue {
	return &SimpleQueue{
		items:       make([]msg.Msg, 0),
		name:        name,
		storage:     storage,
		storageAck:  ackMessageCh,
		storeReader: storeReader,
	}
}
