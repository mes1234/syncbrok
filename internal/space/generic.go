package space

import (
	"github.com/mes1234/syncbrok/internal/msg"
)

type Space interface {
	AddQueue(string)
	Publish(string, msg.Msg)
}
