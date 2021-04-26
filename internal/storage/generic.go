package storage

import (
	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type StorageWriter interface {
	CreateQueue(string) (chan msg.Msg, chan uuid.UUID, FileReader)
	Start()
}
