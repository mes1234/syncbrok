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

func HttpListner(s space.Space, handler msg.Callback) func() {
	queueName := "simpleQueue"
	s.AddQueue(queueName)
	s.Subscribe(queueName, handler)
	homePage := CreateEndpoint(s, queueName)
	return func() {
		http.HandleFunc("/", homePage)
		log.Fatal(http.ListenAndServe(":10000", nil))
	}
}
func CreateEndpoint(s space.Space, queueName string) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		parent := msg.NewSimpleMsg(uuid.Nil, body)
		s.Publish(queueName, parent)
		fmt.Fprintf(w, "Welcome to the HomePage! %v", parent.GetId())
		fmt.Println("Endpoint Hit: homePage")
	}
}
