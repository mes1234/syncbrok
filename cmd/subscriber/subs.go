package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "false")
		return
	}
	fmt.Fprintf(w, "true")
	log.Printf("New msg arrived %v", string(body))

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":20000", nil))
}
