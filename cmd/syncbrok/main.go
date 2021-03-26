package main

import (
	"fmt"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

func main() {

	msg1 := msg.Init(nil, nil)

	msg2 := msg.Init(&msg1, nil)

	q := queue.Create()

	q.Add(msg1)
	q.Add(msg2)

	if msg1 == *msg2.Parent {
		fmt.Printf("msg1 %v ", msg1.Id)
	}

}
