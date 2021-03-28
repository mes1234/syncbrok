package queueService

import (
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/space"
)

type SimpleQueueService struct {
	space.Space
}

func (qs SimpleQueueService) NewQueueHandler(name string) *queue.Queue {
	q := queue.NewSimpleQueue(name)
	qs.Space.AddQueue(q, name)
	return &q
}

func (qs SimpleQueueService) NewMessageHandler(q queue.Queue, m msg.Msg) {
	q.AddMsg(m)
}

func NewSimple(s space.Space) SimpleQueueService {
	return SimpleQueueService{Space: s}
}
