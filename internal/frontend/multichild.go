package frontend

import (
	"github.com/google/uuid"

	"github.com/mes1234/syncbrok/internal/msg"
	"github.com/mes1234/syncbrok/internal/space"
)

func MultiChild(s space.Space, handler func(uuid.UUID)) func() {
	return func() {
		queueName := "simpleQueue"
		s.AddQueue(queueName)
		parent := msg.NewSimpleMsg(nil, nil, handler)
		child1 := msg.NewSimpleMsg(parent, nil, handler)
		child2 := msg.NewSimpleMsg(parent, nil, handler)
		child3 := msg.NewSimpleMsg(parent, nil, handler)
		s.Publish(queueName, parent)
		s.Publish(queueName, child1)
		s.Publish(queueName, child2)
		s.Publish(queueName, child3)

	}
}
