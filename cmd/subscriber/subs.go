package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type simpleMsg struct {
	Id      uuid.UUID `json:"id"`
	Parent  uuid.UUID `json:"parent"`
	Content []byte    `json:"content"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "false")
		return
	}
	msg := simpleMsg{}
	json.Unmarshal(body, &msg)
	content := string(msg.Content)
	if content == "parent" {
		time.Sleep(20 * time.Second)
	}
	fmt.Fprintf(w, "true")
	log.Printf("New msg arrived %v with content %v ", msg, content)

}

func main() {
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":20000", nil)

	client := &http.Client{}
	reqQueue, err := http.NewRequest("GET", "http://localhost:10000/queue", nil)
	if err != nil {
		panic(err)
	}
	reqQueue.Header.Add("queue", "test")
	respQueue, err := client.Do(reqQueue)
	if err != nil {
		panic(err)
	}
	respQueue.Body.Close()

	reqSub, err := http.NewRequest("GET", "http://localhost:10000/subscrib", nil)
	if err != nil {
		panic(err)
	}
	reqSub.Header.Add("queue", "test")
	reqSub.Header.Add("endpoint", "http://localhost:20000")
	respSub, err := client.Do(reqSub)
	if err != nil {
		panic(err)
	}
	respSub.Body.Close()
	parentId := SendMsg("parent", uuid.Nil, client)
	for {
		time.Sleep(200 * time.Millisecond)
		parentId = SendMsg("child", parentId, client)
	}

}

func SendMsg(content string, parentId uuid.UUID, client *http.Client) (msgId uuid.UUID) {
	reqMsg, err := http.NewRequest("GET", "http://localhost:10000/msg", bytes.NewReader([]byte(content)))
	if err != nil {
		panic(err)
	}
	reqMsg.Header.Add("queue", "test")
	if parentId != uuid.Nil {
		reqMsg.Header.Add("ParentId", parentId.String())
	}
	respMsg, err := client.Do(reqMsg)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(respMsg.Body)
	if err != nil {
		panic(err)
	}
	id, err := uuid.Parse(string(body))
	if err != nil {
		panic(err)
	}
	msgId = id
	log.Printf("Posted Msg %v", id)
	respMsg.Body.Close()
	return
}
