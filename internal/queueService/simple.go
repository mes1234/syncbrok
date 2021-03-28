package queueService

import (
	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/queue"
	"github.com/mes1234/syncbrok/internal/space"
)

type SimpleQueueService struct {
}

func (qs SimpleQueueService) NewQueueHandler(name string, s space.Space) *queue.Queue {
	q := queue.NewSimpleQueue(name)
	s.AddQueue(q, name)
	return &q
}

func (qs SimpleQueueService) NewMessageHandler(q queue.Queue, m msg.Msg) {
	q.AddMsg(m)
}

func NewSimple() SimpleQueueService {
	return SimpleQueueService{}
}
