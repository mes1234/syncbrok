package queue

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type Queue interface {
	AddMsg(msg.Msg) uuid.UUID
	FindById(uuid.UUID) (msg.Msg, *sync.WaitGroup)
}
