package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type Queue interface {
	AddMsg(msg.Msg)
	FindById(uuid.UUID) msg.Msg
	AddCallback(msg.Callback, string)
}
