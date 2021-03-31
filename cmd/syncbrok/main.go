package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/frontend"

	"github.com/mes1234/syncbrok/internal/space"
)

func handleMessage(id uuid.UUID) {
	log.Print("This message id is: ", id)
}

func main() {

	// Begining of life
	universe := space.New()
	handler := frontend.MultiChild(universe, handleMessage)
	handler()
	time.Sleep(10 * time.Second)
}
