package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

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
