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

func HttpNewMsgListner(
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
	homePage := CreateNewMsgEndpoint(s, queueName, newMsgCh)
	return func() {
		http.HandleFunc("/", homePage)
		log.Fatal(http.ListenAndServe(":10000", nil))
	}
}
func CreateNewMsgEndpoint(
	s space.Space,
	queueName string,
	newMsgCh chan<- space.Messages) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		parentId := ParseParentId(r)
		newMsg := space.Messages{
			QName:   queueName,
			Content: msg.NewSimpleMsg(parentId, body),
		}
		newMsgCh <- newMsg
		fmt.Fprintf(w, "You posted new Msg with id : %v", newMsg.Content.GetId())
		fmt.Println("New msg arrived")
	}
}

func ParseParentId(r *http.Request) uuid.UUID {
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
