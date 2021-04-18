package space

import (
	"sync"

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
	QName    string
	Handler  msg.Callback
	Endpoint string
}

type Space interface {
	Start(*sync.WaitGroup)
	addQueue(string)
	publish(string, msg.Msg)
	subscribe(string, string)
}
