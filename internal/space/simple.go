package space

import (
	"log"
	"sync"
	"time"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/storage"
)

type SimpleSpace struct {
	queues         map[string]chan<- msg.Msg
	subcribers     map[string]chan<- string
	newMessages    <-chan Message
	newQueues      <-chan Queue
	newSubscribers <-chan Subscriber
	handler        msg.Callback
}

func (s SimpleSpace) Start(wg *sync.WaitGroup) {
	for {
		select {
		case newMsg := <-s.newMessages:
			s.addMsg(newMsg.QName, newMsg.Content)
		case newQueue := <-s.newQueues:
			s.addQueue(newQueue.QName)
		case newSubcriber := <-s.newSubscribers:
			s.addSubscriber(newSubcriber.QName, newSubcriber.Endpoint)
		default:
			time.Sleep(1000)
		}
	}
}

func (s *SimpleSpace) addQueue(queueName string) {

	if _, ok := s.queues[queueName]; !ok {
		store := storage.NewFileWriter()
		storeCh, storeAckCh, storeReader := store.CreateQueue(queueName)
		go store.Start()
		queue := queue.NewSimpleQueue(queueName, storeCh, storeAckCh, storeReader, s.handler)
		s.queues[queueName] = queue.GetMsgCh()
		s.subcribers[queueName] = queue.GetSubscriberCh()
		go queue.Start()

		log.Print("Added new queue ", queueName)
	}
}

func (s SimpleSpace) addMsg(queueName string, m msg.Msg) {
	if _, ok := s.queues[queueName]; ok {
		s.queues[queueName] <- m
		log.Printf("Added new msg to queue %v, with id %v ", queueName, m.GetId())
	}
}

func (s SimpleSpace) addSubscriber(queueName string, endpoint string) {
	if _, ok := s.subcribers[queueName]; ok {
		s.subcribers[queueName] <- endpoint
		log.Printf("Added new msg handler to queue %v ", queueName)
	}

}
func New(handler msg.Callback) (Space, chan<- Message, chan<- Queue, chan<- Subscriber) {
	newMessagesCh := make(chan Message)
	newQueuesCh := make(chan Queue)
	newSubscribersCh := make(chan Subscriber)
	simpleSpace := &SimpleSpace{
		queues:         make(map[string]chan<- msg.Msg),
		subcribers:     make(map[string]chan<- string),
		newMessages:    newMessagesCh,
		newQueues:      newQueuesCh,
		newSubscribers: newSubscribersCh,
		handler:        handler,
	}
	return simpleSpace, newMessagesCh, newQueuesCh, newSubscribersCh
}
