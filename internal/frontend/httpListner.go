package frontend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/space"
)

func HttpNewMsgController(
	handler msg.Callback,
	newMsgCh chan<- space.Messages,
	newSubscribersCh chan<- space.Subscribers) {
	queueName := "simpleQueue"
	newSubscribersCh <- space.Subscribers{
		QName:   queueName,
		Handler: handler,
	}
	msgHandler := createNewMsgEndpoint(newMsgCh)
	http.HandleFunc("/msg", msgHandler)
}

func HttpNewQueueController(newQueueCh chan<- space.Queues) {
	queueHandler := createNewQueueEndpoint(newQueueCh)
	http.HandleFunc("/queue", queueHandler)
}

func HttpStart() {
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func createNewQueueEndpoint(newQueueCh chan<- space.Queues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queueName := parseQueueName(r)
		if queueName == "" {
			log.Printf("New queue requested but no name provided")
			fmt.Fprintf(w, "You posted empty queue name, queue name shall be provided")
			return
		}
		newQueue := space.Queues{
			QName: queueName,
		}
		newQueueCh <- newQueue
		fmt.Fprintf(w, "You posted new queue with name : %v", queueName)
		log.Printf("New queue request arrived")
	}
}

func createNewMsgEndpoint(newMsgCh chan<- space.Messages) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		parentId := parseParentId(r)
		queueName := parseQueueName(r)
		newMsg := space.Messages{
			QName:   queueName,
			Content: msg.NewSimpleMsg(parentId, body),
		}
		newMsgCh <- newMsg
		fmt.Fprintf(w, "You posted new Msg with id : %v", newMsg.Content.GetId())
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
