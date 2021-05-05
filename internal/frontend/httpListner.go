package frontend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/space"
)

func httpNewMsgController(newMsgCh chan<- space.Message) {
	msgHandler := createNewMsgEndpoint(newMsgCh)
	http.HandleFunc("/msg", msgHandler)
}

func httpNewSubscriberController(newSubscribersCh chan<- space.Subscriber) {
	subscriberHandler := createNewSubscriberEndpoint(newSubscribersCh)
	http.HandleFunc("/subscrib", subscriberHandler)
}

func httpNewQueueController(newQueueCh chan<- space.Queue) {
	queueHandler := createNewQueueEndpoint(newQueueCh)
	http.HandleFunc("/queue", queueHandler)
}

func HttpStart(
	wg *sync.WaitGroup,
	newMsgCh chan<- space.Message,
	newSubscribersCh chan<- space.Subscriber,
	newQueueCh chan<- space.Queue) {
	httpNewMsgController(newMsgCh)
	httpNewSubscriberController(newSubscribersCh)
	httpNewQueueController(newQueueCh)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func createNewSubscriberEndpoint(newSubscribersCh chan<- space.Subscriber) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queueName := parseQueueName(r)
		endpointName := parseEndpointName(r)
		newSubscribersCh <- space.Subscriber{
			QName:    queueName,
			Endpoint: endpointName,
		}
		fmt.Fprintf(w, "You posted new subscriber for queue  : %v", queueName)
		log.Printf("New subscriber request arrived")
	}
}

func createNewQueueEndpoint(newQueueCh chan<- space.Queue) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queueName := parseQueueName(r)
		if queueName == "" {
			log.Printf("New queue requested but no name provided")
			fmt.Fprintf(w, "You posted empty queue name, queue name shall be provided")
			return
		}
		newQueue := space.Queue{
			QName: queueName,
		}
		newQueueCh <- newQueue
		fmt.Fprintf(w, "You posted new queue with name : %v", queueName)
		log.Printf("New queue request arrived")
	}
}

func createNewMsgEndpoint(newMsgCh chan<- space.Message) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		parentId := parseParentId(r)
		queueName := parseQueueName(r)
		newMsg := space.Message{
			QName:   queueName,
			Content: msg.NewSimpleMsg(parentId, body),
		}
		newMsgCh <- newMsg
		fmt.Fprintf(w, "%v", newMsg.Content.GetId())
		log.Printf("New msg arrived")
	}
}

func parseQueueName(r *http.Request) string {
	queueName := r.Header.Get("queue")
	if queueName == "" {
		return ""
	} else {
		return queueName
	}
}

func parseEndpointName(r *http.Request) string {
	endpointName := r.Header.Get("endpoint")
	if endpointName == "" {
		return ""
	} else {
		return endpointName
	}
}

func parseParentId(r *http.Request) uuid.UUID {
	parentIdStr := r.Header.Get("ParentId")
	if parentIdStr == "" {
		return uuid.Nil
	} else {
		parentId, err := uuid.Parse(parentIdStr)
		if err != nil {
			return uuid.Nil
		} else {
			return parentId
		}
	}
}
