package main

import (
	"github.com/mes1234/syncbrok/internal/frontend"
	"github.com/mes1234/syncbrok/internal/queueService"

	"github.com/mes1234/syncbrok/internal/space"
)

func main() {

	// Begining of life
	universe := space.New()
	queueListner := queueService.NewSimple(universe)
	handler := frontend.Fake(queueListner)
	handler()

}
