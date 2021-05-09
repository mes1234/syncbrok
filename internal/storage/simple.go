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

func (fw *FileWriter) CreateQueue(queueName string) (addMsgCh chan msg.MsgContent, contentReader FileReader) {
	fileContent, err := os.Create(fw.path + queueName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fw.fileContent = bufio.NewReadWriter(bufio.NewReader(fileContent), bufio.NewWriter(fileContent))

	newMessagesCh := make(chan msg.MsgContent, 100)
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

func (fw *FileWriter) addToStore(m msg.MsgContent) {
	content := m.Content
	msg := MsgSave{
		StartOffset: fw.offset,
		Len:         len(content),
		Id:          m.Id,
	}
	fw.lookup[m.Id] = msg

	fw.offset = fw.offset + int64(len(content))
	fw.fileContent.Write(content)
	fw.fileContent.Flush()
}

func prepareReader(file *os.File, msgLocation map[uuid.UUID]MsgSave) FileReader {
	fname := file.Name()
	return func(u uuid.UUID) []byte {
		log.Printf("attempt to retrieve content of %v", u)
		f, err := os.Open(fname)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		offset := (msgLocation[u])
		buf := make([]byte, offset.Len)

		f.ReadAt(buf, int64(offset.StartOffset))
		strBuf := string(buf)
		log.Print("will send", strBuf)
		return buf
	}
}
