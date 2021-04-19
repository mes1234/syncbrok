package storage

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"time"

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
	fileIndex   *bufio.Writer
	offset      int64
	addMsgCh    <-chan msg.Msg
	lookup      map[uuid.UUID]MsgSave
	buffer      bytes.Buffer
}

type FileReader func(uuid.UUID) []byte

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
	msg := MsgSave{
		StartOffset: fw.offset,
		Len:         len(content),
		Id:          m.GetId(),
		Parent:      m.GetParentId(),
	}
	fw.lookup[m.GetId()] = msg
	encoder := gob.NewEncoder(&fw.buffer)
	err := encoder.Encode(msg)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	log.Printf("bytes %v", fw.buffer.Bytes())
	fw.fileIndex.Write(fw.buffer.Bytes())
	fw.buffer.Reset()
	fw.fileIndex.Flush()

	fw.offset = fw.offset + int64(len(content))
	fw.fileContent.Write(content)
	fw.fileContent.Flush()
}

func (fw *FileWriter) CreateQueue(queueName string) (addMsgCh chan<- msg.Msg, reader FileReader) {
	fileContent, err := os.Create(fw.path + queueName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fileIndex, err := os.Create(fw.path + queueName + "_id")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fw.fileContent = bufio.NewWriter(fileContent)
	fw.fileIndex = bufio.NewWriter(fileIndex)
	ch := make(chan msg.Msg)
	addMsgCh = ch
	fw.addMsgCh = ch
	reader = prepareReader(fileContent, fw.lookup)
	return
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

func NewFileWriter() StorageWriter {
	fw := FileWriter{
		lookup: make(map[uuid.UUID]MsgSave),
		path:   "C:\\Users\\witol\\go\\syncbrok\\temp\\",
		buffer: bytes.Buffer{},
	}
	return &fw
}
