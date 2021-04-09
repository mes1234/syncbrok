package storage

import (
	"bufio"
	"io/ioutil"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type MsgSave struct {
	startOffset int64
	len         int
}

type FileWriter struct {
	file     *bufio.Writer
	offset   int64
	addMsgCh <-chan msg.Msg
	lookup   map[uuid.UUID]MsgSave
}

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

func (fw *FileWriter) addToStore(m msg.Msg) {
	content := m.GetContent()
	fw.lookup[m.GetId()] = MsgSave{
		startOffset: fw.offset,
		len:         len(content),
	}
	fw.offset = fw.offset + int64(len(content))
	fw.file.Write(content)
	fw.file.Flush()
}

func (fw *FileWriter) CreateQueue(queueName string) chan<- msg.Msg {
	file, err := ioutil.TempFile("C:\\Users\\witol\\go\\syncbrok\\temp", queueName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fw.file = bufio.NewWriter(file)
	addMsgCh := make(chan msg.Msg)
	fw.addMsgCh = addMsgCh
	return addMsgCh
}

func NewFileWriter() StorageWriter {
	return &FileWriter{
		lookup: make(map[uuid.UUID]MsgSave),
	}
}
