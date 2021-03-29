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
		topLevelMsg := msg.NewSimpleMsg(nil, nil, handler)
		qs.NewMessageHandler(*q, topLevelMsg)
		var parent msg.Msg = nil
		var child msg.Msg = nil

		parent = msg.NewSimpleMsg(topLevelMsg, nil, handler)
		qs.NewMessageHandler(*q, parent)
		for {

			child = msg.NewSimpleMsg(parent, nil, handler)
			qs.NewMessageHandler(*q, child)
			parent = msg.NewSimpleMsg(child, nil, handler)
			qs.NewMessageHandler(*q, parent)
			time.Sleep(1 * time.Second)
		}
	}
}
