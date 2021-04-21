package storage

import (
	"bufio"
	"bytes"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type MsgSave struct {
	StartOffset int64
	Len         int
	Id          uuid.UUID
	Parent      uuid.UUID
}

type FileWriter struct {
	path        string
	fileContent *bufio.Writer
	fileIndex   *bufio.ReadWriter
	offset      int64
	addMsgCh    <-chan msg.Msg
	msgAckCh    <-chan uuid.UUID
	lookup      map[uuid.UUID]MsgSave
	buffer      bytes.Buffer
}

type FileReader func(uuid.UUID) []byte
