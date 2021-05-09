package storage

import (
	"bufio"
	"bytes"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type StorageWriter interface {
	CreateQueue(string) (chan msg.MsgContent, FileReader)
	Start()
}

type MsgSave struct {
	StartOffset int64
	Len         int
	Id          uuid.UUID
}

type FileWriter struct {
	path        string
	fileContent *bufio.ReadWriter
	offset      int64
	addMsgCh    <-chan msg.MsgContent
	lookup      map[uuid.UUID]MsgSave
	buffer      bytes.Buffer
}

type FileReader func(uuid.UUID) []byte
