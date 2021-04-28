package queue

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/storage"
)

type SimpleQueue struct {
	items           []msg.Msg
	deliveredItems  []uuid.UUID
	name            string
	subscribers     []string
	handler         msg.Callback
	storage         chan<- msg.Msg
	storageAck      chan<- uuid.UUID
	storeReader     storage.FileReader
	newMsgCh        chan msg.Msg
	newSubscriberCh chan string
}

func (q SimpleQueue) findById(id uuid.UUID) msg.Msg {
	if true == true {
		for _, element := range q.items {
			if element.GetId() == id {
				return element
			}
		}
		return nil
	}
	return nil
}

func (q *SimpleQueue) GetSubscriberCh() chan<- string {
	return q.newSubscriberCh
}

func (q *SimpleQueue) GetMsgCh() chan<- msg.Msg {
	return q.newMsgCh
}

func (q *SimpleQueue) Start() {
	for {
		select {
		case newMsg := <-q.newMsgCh:
			q.addMsg(newMsg)
		case newSubscriber := <-q.newSubscriberCh:
			q.addCallback(newSubscriber)
		default:
			time.Sleep(1000)
		}
	}
}

//Add item to end of queue
func (q *SimpleQueue) addMsg(m msg.Msg) {

	q.storage <- m
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)

	parentId := m.GetParentId()
	var parent msg.Msg = nil
	if parentId != uuid.Nil || q.delivered(parentId) {
		parent = q.findById(parentId)
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

func (q *SimpleQueue) addCallback(endpoint string) {
	q.subscribers = append(q.subscribers, endpoint)
}

func (q *SimpleQueue) captureDelivery(ackMessageCh chan<- uuid.UUID, proxyMsgAckCh <-chan uuid.UUID) {
	for {
		msgAck := <-proxyMsgAckCh
		q.deliveredItems = append(q.deliveredItems, msgAck)
		ackMessageCh <- msgAck
	}

}

func NewSimpleQueue(name string, storage chan msg.Msg, ackMessageCh chan uuid.UUID, storeReader storage.FileReader, handler msg.Callback) Queue {
	proxyMsgAckCh := make(chan uuid.UUID)
	q := SimpleQueue{
		items:           make([]msg.Msg, 0),
		name:            name,
		storage:         storage,
		storageAck:      proxyMsgAckCh,
		storeReader:     storeReader,
		deliveredItems:  make([]uuid.UUID, 0),
		newMsgCh:        make(chan msg.Msg, 100),
		newSubscriberCh: make(chan string, 100),
		handler:         handler,
	}

	go q.captureDelivery(ackMessageCh, proxyMsgAckCh)
	return &q

}
