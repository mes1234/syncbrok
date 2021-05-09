package space

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/storage"
)

type HandlerFactoryFunc func(storage.FileReader) func(uuid.UUID, string, *sync.WaitGroup)

type SimpleSpace struct {
	queues         map[string]chan<- msg.Msg
	subcribers     map[string]chan<- string
	newMessages    <-chan Message
	newQueues      <-chan Queue
	newSubscribers <-chan Subscriber
	handler        HandlerFactoryFunc
}

func (s SimpleSpace) Start(wg *sync.WaitGroup) {
	for {
		select {
		case newMsg := <-s.newMessages:
			s.addMsg(newMsg.QName, newMsg.Content)
		case newQueue := <-s.newQueues:
			s.addQueue(newQueue.QName, newQueue.Storage)
		case newSubcriber := <-s.newSubscribers:
			s.addSubscriber(newSubcriber.QName, newSubcriber.Endpoint)
		default:
			time.Sleep(1000)
		}
	}
}

func (s *SimpleSpace) addStorageToQueue(queueName string, storagePath string) (storeCh chan msg.MsgContent, storeReader storage.FileReader) {
	store := storage.NewFileWriter(storagePath)
	storeCh, storeReader = store.CreateQueue(queueName)
	go store.Start()
	return
}

func (s *SimpleSpace) initNewQueue(queueName string, storeCh chan msg.MsgContent, storeReader storage.FileReader) {
	queue := queue.NewSimpleQueue(queueName, storeCh, s.handler(storeReader))
	s.queues[queueName] = queue.GetMsgCh()
	s.subcribers[queueName] = queue.GetSubscriberCh()
	go queue.Start()
}

func (s *SimpleSpace) addQueue(queueName string, storagePath string) {

	if _, ok := s.queues[queueName]; !ok {
		storeCh, storeReader := s.addStorageToQueue(queueName, storagePath)
		s.initNewQueue(queueName, storeCh, storeReader)
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
func New(handler HandlerFactoryFunc) (Space, chan<- Message, chan<- Queue, chan<- Subscriber) {
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
