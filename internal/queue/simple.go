package queue

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type SimpleQueue struct {
	items           []msg.Msg
	deliveredItems  []uuid.UUID
	name            string
	subscribers     []string
	handler         msg.Callback
	storage         chan<- msg.MsgContent
	newMsgCh        chan msg.Msg
	newSubscriberCh chan string
}

func (q SimpleQueue) findById(id uuid.UUID) msg.Msg {
	for _, element := range q.items {
		if element.GetId() == id {
			return element
		}
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

	q.storage <- msg.MsgContent{
		Id:      m.GetId(),
		Content: m.GetContent(),
	}
	m.RemovePayload()
	q.items = append(q.items, m)
	log.Print("Added item to  queue :", q.name)

	parentId := m.GetParentId()
	var parent msg.Msg = nil

	if parentId != uuid.Nil || q.delivered(parentId) {
		parent = q.findById(parentId)
		go m.Process(parent.GetWaiter(), q.handler, q.subscribers)
	} else {
		go m.Process(nil, q.handler, q.subscribers)
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

func NewSimpleQueue(name string, storage chan msg.MsgContent, handler msg.Callback) Queue {

	return &SimpleQueue{
		items:           make([]msg.Msg, 0),
		name:            name,
		storage:         storage,
		deliveredItems:  make([]uuid.UUID, 0),
		newMsgCh:        make(chan msg.Msg, 100),
		newSubscriberCh: make(chan string, 100),
		handler:         handler,
	}

}
