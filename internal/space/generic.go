package space

import (
	"sync"

	"github.com/mes1234/syncbrok/internal/msg"
)

type Message struct {
	QName   string
	Content msg.Msg
}

type Queue struct {
	QName string
}

type Subscriber struct {
	QName    string
	Handler  msg.Callback
	Endpoint string
}

type Space interface {
	Start(*sync.WaitGroup)
	addQueue(string)
	addMsg(string, msg.Msg)
	addSubscriber(string, string)
}
