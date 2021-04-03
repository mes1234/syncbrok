package space

import (
	"log"
	"time"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

type SimpleSpace struct {
	queues         map[string]queue.Queue
	newMessages    <-chan Messages
	newQueues      <-chan Queues
	newSubscribers <-chan Subscribers
}

func (s SimpleSpace) Start() {
	for {
		select {
		case newMsg := <-s.newMessages:
			s.publish(newMsg.QName, newMsg.Content)
		case newQueue := <-s.newQueues:
			s.addQueue(newQueue.QName)
		case newSubcriber := <-s.newSubscribers:
			s.subscribe(newSubcriber.QName, newSubcriber.Handler)
		default:
			time.Sleep(1000)
		}
	}
}

func (s SimpleSpace) addQueue(name string) {
	log.Print("Added new queue ", name)
	s.queues[name] = queue.NewSimpleQueue(name)
}

func (s SimpleSpace) publish(queueName string, m msg.Msg) {
	s.queues[queueName].AddMsg(m)
}

func (s SimpleSpace) subscribe(queueName string, callback msg.Callback) {
	s.queues[queueName].AddCallback(callback)
}

func New() (Space, chan<- Messages, chan<- Queues, chan<- Subscribers) {
	newMessagesCh := make(chan Messages)
	newQueuesCh := make(chan Queues)
	newSubscribersCh := make(chan Subscribers)
	simpleSpace := SimpleSpace{
		queues:         make(map[string]queue.Queue),
		newMessages:    newMessagesCh,
		newQueues:      newQueuesCh,
		newSubscribers: newSubscribersCh,
	}
	return simpleSpace, newMessagesCh, newQueuesCh, newSubscribersCh
}
