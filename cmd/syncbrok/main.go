package main

import (
	"fmt"

	"github.com/mes1234/syncbrok/internal/msg"
)

func main() {

	msg1 := msg.Init(nil, nil)

	msg2 := msg.Init(&msg1, nil)

	if msg1 == *msg2.Parent {
		fmt.Printf("msg1 %v ", msg1.Id)
	}

}
