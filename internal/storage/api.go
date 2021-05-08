package storage

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

func (fw FileWriter) Start() {
	for {
		select {
		case newMsg := <-fw.addMsgCh:
			fw.addToStore(newMsg)
		default:
			time.Sleep(1000)
		}

	}
}

func (fw *FileWriter) CreateQueue(queueName string) (addMsgCh chan msg.Msg, contentReader FileReader) {
	fileContent, err := os.Create(fw.path + queueName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fw.fileContent = bufio.NewReadWriter(bufio.NewReader(fileContent), bufio.NewWriter(fileContent))

	newMessagesCh := make(chan msg.Msg, 100)
	addMsgCh = newMessagesCh
	fw.addMsgCh = newMessagesCh

	contentReader = prepareReader(fileContent, fw.lookup)
	return
}

func NewFileWriter(path string) StorageWriter {
	fw := FileWriter{
		lookup: make(map[uuid.UUID]MsgSave),
		path:   path,
		buffer: bytes.Buffer{},
	}
	return &fw
}
