package main

import (
	"fmt"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

func main() {

	msg1 := msg.New(nil, nil)

	msg2 := msg.New(&msg1, nil)

	q := queue.New()

	q.Add(msg1)
	q.Add(msg2)

	fmt.Printf("msg1 %v ", msg1.GetId())

}
