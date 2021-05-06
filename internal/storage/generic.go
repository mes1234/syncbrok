package storage

import (
	"github.com/mes1234/syncbrok/internal/msg"
)

type StorageWriter interface {
	CreateQueue(string) (chan msg.Msg, FileReader)
	Start()
}
