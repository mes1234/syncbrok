package frontend

import (
	"time"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queueService"
)

func Fake(qs queueService.QueueService) func() {
	return func() {
		q := qs.NewQueueHandler("simpleQueue")
		for {
			qs.NewMessageHandler(*q, msg.NewSimpleMsg(nil, nil))
			time.Sleep(2 * time.Second)
		}
	}
}
