package space

import (
	"github.com/mes1234/syncbrok/internal/msg"
)

type Messages struct {
	QName   string
	Content msg.Msg
}

type Queues struct {
	QName string
}

type Subscribers struct {
	QName   string
	Handler msg.Callback
}

type Space interface {
	Start()
	addQueue(string)
	publish(string, msg.Msg)
	subscribe(string, msg.Callback)
}
