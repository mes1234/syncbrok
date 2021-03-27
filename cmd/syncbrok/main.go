package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mes1234/syncbrok/internal/queueService"

	"github.com/mes1234/syncbrok/internal/space"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

func main() {

	// Begining of life
	universe := space.NewSpace()
	queueListner := queueService.NewSimpleQueueService()
	queueListner.Start(universe)
	log.Print("GO ON")
	time.Sleep(10 * time.Second)
	universe.AddQueue(queue.NewSimpleQueue(), "first")

	msg1 := msg.NewSimpleMsg(nil, nil)

	msg2 := msg.NewSimpleMsg(&msg1, nil)

	q := queue.NewSimpleQueue()

	q.Add(msg1)
	q.Add(msg2)

	fmt.Printf("msg1 %v ", msg1.GetId())

}
