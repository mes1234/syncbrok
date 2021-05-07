package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/storage"
)

func HttphandleMessageFactory(storageReader storage.FileReader) func(uuid.UUID, string, *sync.WaitGroup) {

	return func(msgId uuid.UUID, endpoint string, callbackWg *sync.WaitGroup) {
		for {
			content := storageReader(msgId)
			bodyReq := bytes.NewBuffer(content)
			resp, err := http.Post(endpoint, "application/json", bodyReq)
			if err != nil {
				log.Printf("An Error Occured %v", err)
				continue
			}
			defer resp.Body.Close()
			//Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("error fetching response %v", err)
				resp.Body.Close()
				continue
			}
			sb := string(body)
			if sb == "true" {
				callbackWg.Done()
				break
			}
		}
	}
}
