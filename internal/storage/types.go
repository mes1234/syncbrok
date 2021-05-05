package storage

import (
	"bufio"
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type MsgSave struct {
	StartOffset int64
	Len         int
	Id          uuid.UUID
	Parent      uuid.UUID
	TimeStamp   time.Time
}

type FileWriter struct {
	path        string
	fileContent *bufio.ReadWriter
	fileIndex   *bufio.ReadWriter
	fileAck     *bufio.ReadWriter
	offset      int64
	addMsgCh    <-chan msg.Msg
	msgAckCh    <-chan uuid.UUID
	lookup      map[uuid.UUID]MsgSave
	buffer      bytes.Buffer
}

type FileReader func(uuid.UUID) []byte
