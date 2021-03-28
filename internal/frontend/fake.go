package frontend

import (
	"time"

	"github.com/google/uuid"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queueService"
)

func Fake(qs queueService.QueueService, handler func(uuid.UUID)) func() {
	return func() {
		q := qs.NewQueueHandler("simpleQueue")
		for {
			qs.NewMessageHandler(*q, msg.NewSimpleMsg(nil, nil, handler))
			time.Sleep(2 * time.Second)
		}
	}
}
