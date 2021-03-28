package frontend

import (
	"time"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queueService"
	"github.com/mes1234/syncbrok/internal/space"
)

func Fake(qs queueService.QueueService, s space.Space) func() {
	return func() {
		q := qs.NewQueueHandler("simpleQueue", s)
		for {
			qs.NewMessageHandler(*q, msg.NewSimpleMsg(nil, nil))
			time.Sleep(2 * time.Second)
		}
	}
}
