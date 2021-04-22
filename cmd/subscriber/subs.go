package main

import (
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
		time.Sleep(40 * time.Second)
	}
	fmt.Fprintf(w, "true")
	log.Printf("New msg arrived %v with content %v ", msg, content)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":20000", nil)

}
