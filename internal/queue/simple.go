package queue

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/storage"
)

type SimpleQueue struct {
	items          []msg.Msg
	deliveredItems []uuid.UUID
	name           string
	subscribers    []string
	handler        msg.Callback
	storage        chan<- msg.Msg
	storageAck     chan<- uuid.UUID
	storeReader    storage.FileReader
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
	if parentId != uuid.Nil || q.delivered(parentId) {
		parent = q.FindById(parentId)
		go m.Process(parent.GetWaiter(), q.handler, q.subscribers, q.storageAck)
	} else {
		go m.Process(nil, q.handler, q.subscribers, q.storageAck)
	}

}

//If msg is deliver return true
//otherwise return false
func (q SimpleQueue) delivered(u uuid.UUID) bool {
	for _, item := range q.deliveredItems {
		if item == u {
			return true
		}
	}
	return false
}

func (q *SimpleQueue) AddCallback(callback msg.Callback, endpoint string) {
	q.subscribers = append(q.subscribers, endpoint)
	q.handler = callback
}

func (q *SimpleQueue) captureDelivery(ackMessageCh chan<- uuid.UUID, proxyMsgAckCh <-chan uuid.UUID) {
	for {
		msgAck := <-proxyMsgAckCh
		q.deliveredItems = append(q.deliveredItems, msgAck)
		ackMessageCh <- msgAck
	}

}

func NewSimpleQueue(name string, storage chan msg.Msg, ackMessageCh chan uuid.UUID, storeReader storage.FileReader) Queue {
	proxyMsgAckCh := make(chan uuid.UUID)
	q := SimpleQueue{
		items:          make([]msg.Msg, 0),
		name:           name,
		storage:        storage,
		storageAck:     proxyMsgAckCh,
		storeReader:    storeReader,
		deliveredItems: make([]uuid.UUID, 0),
	}

	go q.captureDelivery(ackMessageCh, proxyMsgAckCh)
	return &q

}
