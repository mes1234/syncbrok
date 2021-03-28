package queueService

import (
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
)

type QueueService interface {
	NewQueueHandler(string) *queue.Queue
	NewMessageHandler(queue.Queue, msg.Msg)
}
