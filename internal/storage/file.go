package storage

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

type MsgSave struct {
	startOffset int64
	len         int
}

type FileWriter struct {
	path     string
	file     *bufio.Writer
	offset   int64
	addMsgCh <-chan msg.Msg
	lookup   map[uuid.UUID]MsgSave
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
	fw.lookup[m.GetId()] = MsgSave{
		startOffset: fw.offset,
		len:         len(content),
	}
	fw.offset = fw.offset + int64(len(content))
	fw.file.Write(content)
	fw.file.Flush()
}

func (fw *FileWriter) CreateQueue(queueName string) (addMsgCh chan<- msg.Msg, reader FileReader) {
	file, err := os.Create(fw.path + queueName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fw.file = bufio.NewWriter(file)
	ch := make(chan msg.Msg)
	addMsgCh = ch
	fw.addMsgCh = ch
	reader = prepareReader(file, fw.lookup)
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
		buf := make([]byte, offset.len)

		f.ReadAt(buf, int64(offset.startOffset))
		strBuf := string(buf)
		log.Print("will send", strBuf)
		return buf
	}

}

func NewFileWriter() StorageWriter {
	return &FileWriter{
		lookup: make(map[uuid.UUID]MsgSave),
		path:   "C:\\Users\\witol\\go\\syncbrok\\temp\\",
	}
}
