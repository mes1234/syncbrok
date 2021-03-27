package queue

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type Queue interface {
	Add(msg.IMsg) uuid.UUID
}
