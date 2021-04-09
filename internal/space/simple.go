package space

import (
	"log"
	"time"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/storage"
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
			s.subscribe(newSubcriber.QName, newSubcriber.Handler, newSubcriber.Endpoint)
		default:
			time.Sleep(1000)
		}
	}
}

func (s SimpleSpace) addQueue(queueName string) {

	if _, ok := s.queues[queueName]; !ok {
		store := storage.NewFileWriter()
		storeCh := store.CreateQueue(queueName)
		go store.Start()
		s.queues[queueName] = queue.NewSimpleQueue(queueName, storeCh)
		log.Print("Added new queue ", queueName)
	}
}

func (s SimpleSpace) publish(queueName string, m msg.Msg) {
	if _, ok := s.queues[queueName]; ok {
		s.queues[queueName].AddMsg(m)
		log.Printf("Added new msg to queue %v, with id %v ", queueName, m.GetId())
	}
}

func (s SimpleSpace) subscribe(queueName string, callback msg.Callback, endpoint string) {
	if _, ok := s.queues[queueName]; ok {
		s.queues[queueName].AddCallback(callback, endpoint)
		log.Printf("Added new msg handler to queue %v ", queueName)
	}

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
