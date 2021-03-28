package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/frontend"
	"github.com/mes1234/syncbrok/internal/queueService"

	"github.com/mes1234/syncbrok/internal/space"
)

func handleMessage(id uuid.UUID) {
	log.Print("This message id is: ", id)
}

func main() {

	// Begining of life
	universe := space.New()
	queueListner := queueService.NewSimple(universe)
	handler := frontend.Fake(queueListner, handleMessage)
	handler()

}
