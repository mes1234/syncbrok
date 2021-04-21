package storage

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
	fileIndex   *bufio.ReadWriter
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

func (fw *FileWriter) decodeMsgSave(buffer []byte) {
	gobReader := bytes.NewReader(buffer)
	decoder := gob.NewDecoder(gobReader)
	newMsg := MsgSave{}
	err := decoder.Decode(&newMsg)
	if err != nil {
		panic(err)
	}
	fw.lookup[newMsg.Id] = newMsg
	fw.offset = fw.offset + int64(newMsg.Len)
}

func (fw *FileWriter) recoverMsges(queueName string) {

	reader := fw.fileIndex.Reader
	for {

		buffer := make([]byte, binary.MaxVarintLen64)
		binary.Read(reader, binary.LittleEndian, buffer)
		count, err := binary.ReadUvarint(bytes.NewReader(buffer))
		if err != nil || count == 0 {
			break
		}

		buffer = make([]byte, count)
		err = binary.Read(reader, binary.LittleEndian, buffer)
		if err != nil {
			break
		}

		go fw.decodeMsgSave(buffer)

	}
}

func getBytes(i int) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(i))
	return buf
}

func (fw *FileWriter) EncodeMsg(msg MsgSave) (bytes []byte) {
	encoder := gob.NewEncoder(&fw.buffer)
	err := encoder.Encode(msg)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	bytes = fw.buffer.Bytes()
	fw.buffer.Reset()
	return
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

	msgBytes := fw.EncodeMsg(msg)
	fw.fileIndex.Write(getBytes(len(msgBytes))) // save size of gob
	fw.fileIndex.Write(msgBytes)                //save gob

	fw.fileIndex.Flush()

	fw.offset = fw.offset + int64(len(content))
	fw.fileContent.Write(content)
	fw.fileContent.Flush()
}

func (fw *FileWriter) CreateQueue(queueName string) (addMsgCh chan<- msg.Msg, reader FileReader) {
	fileContent, err := os.OpenFile(fw.path+queueName, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fileIndex, err := os.OpenFile(fw.path+queueName+"_id", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fw.fileContent = bufio.NewWriter(fileContent)
	fw.fileIndex = bufio.NewReadWriter(bufio.NewReader(fileIndex), bufio.NewWriter(fileIndex))

	fw.recoverMsges(queueName)
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
