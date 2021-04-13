package frontend

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/space"
)

func httpNewMsgController(newMsgCh chan<- space.Messages) {
	msgHandler := createNewMsgEndpoint(newMsgCh)
	http.HandleFunc("/msg", msgHandler)
}

func httpNewSubscriberController(
	handler msg.Callback,
	newSubscribersCh chan<- space.Subscribers) {
	subscriberHandler := createNewSubscriberEndpoint(newSubscribersCh, handler)
	http.HandleFunc("/subscrib", subscriberHandler)
}

func httpNewQueueController(newQueueCh chan<- space.Queues) {
	queueHandler := createNewQueueEndpoint(newQueueCh)
	http.HandleFunc("/queue", queueHandler)
}

func HttpStart(
	newMsgCh chan<- space.Messages,
	newSubscribersCh chan<- space.Subscribers,
	newQueueCh chan<- space.Queues,
	handler msg.Callback) {
	httpNewMsgController(newMsgCh)
	httpNewSubscriberController(handler, newSubscribersCh)
	httpNewQueueController(newQueueCh)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func createNewSubscriberEndpoint(newSubscribersCh chan<- space.Subscribers, handler msg.Callback) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queueName := parseQueueName(r)
		endpointName := parseEndpointName(r)
		newSubscribersCh <- space.Subscribers{
			QName:    queueName,
			Handler:  handler,
			Endpoint: endpointName,
		}
		fmt.Fprintf(w, "You posted new subscriber for queue  : %v", queueName)
		log.Printf("New subscriber request arrived")
	}
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

func HttphandleMessage(content []byte, endpoint string) bool {
	bodyReq := bytes.NewBuffer(content)
	resp, err := http.Post(endpoint, "application/json", bodyReq)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return false
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	if sb == "true" {
		return true
	} else {
		return false
	}
}
