package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func HttphandleMessage(content []byte, endpoint string, callbackWg *sync.WaitGroup) {
	bodyReq := bytes.NewBuffer(content)
	resp, err := http.Post(endpoint, "application/json", bodyReq)
	if err != nil {
		log.Printf("An Error Occured %v", err)
		time.Sleep(time.Second * 3)
		go HttphandleMessage(content, endpoint, callbackWg)
		return
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error fetching response %v", err)
		go HttphandleMessage(content, endpoint, callbackWg)
		return
	}
	sb := string(body)
	if sb == "true" {
		callbackWg.Done()
	} else {
		go HttphandleMessage(content, endpoint, callbackWg)
		return
	}
}
