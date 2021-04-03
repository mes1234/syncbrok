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

func HttpListner(
	handler msg.Callback,
	s space.Space,
	newMsgCh chan<- space.Messages,
	newQueueCh chan<- space.Queues,
	newSubscribersCh chan<- space.Subscribers) func() {
	queueName := "simpleQueue"
	newQueueCh <- space.Queues{
		QName: queueName,
	}
	newSubscribersCh <- space.Subscribers{
		QName:   queueName,
		Handler: handler,
	}
	homePage := CreateEndpoint(s, queueName, newMsgCh)
	return func() {
		http.HandleFunc("/", homePage)
		log.Fatal(http.ListenAndServe(":10000", nil))
	}
}
func CreateEndpoint(s space.Space, queueName string, newMsgCh chan<- space.Messages) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		newMsg := space.Messages{
			QName:   queueName,
			Content: msg.NewSimpleMsg(uuid.Nil, body),
		}
		newMsgCh <- newMsg
		fmt.Fprintf(w, "Welcome to the HomePage! %v", newMsg.Content.GetId())
		fmt.Println("Endpoint Hit: homePage")
	}
}
