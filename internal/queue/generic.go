package queue

import (
	"github.com/mes1234/syncbrok/internal/msg"
)

type Queue interface {
	Start()
	GetMsgCh() chan<- msg.Msg
	GetSubscriberCh() chan<- string
}
