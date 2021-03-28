package queueService

import (
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/space"
)

type QueueService interface {
	NewQueueHandler(string, space.Space) *queue.Queue
	NewMessageHandler(queue.Queue, msg.Msg)
}
