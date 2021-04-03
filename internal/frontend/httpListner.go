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

func HttpListner(s space.Space, handler msg.Callback, newMsgCh chan<- space.Messages, newQueuCh chan<- space.Queues) func() {
	queueName := "simpleQueue"
	newQueuCh <- space.Queues{
		QName: queueName,
	}
	s.Subscribe(queueName, handler)
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
