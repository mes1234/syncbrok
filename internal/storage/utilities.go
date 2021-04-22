package storage

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

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

		fw.decodeMsgSave(buffer)

	}
}

func getBytes(i int) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(i))
	return buf
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

	msgBytes := fw.encodeMsgSave(msg)
	fw.fileIndex.Write(getBytes(len(msgBytes))) // save size of gob
	fw.fileIndex.Write(msgBytes)                //save gob

	fw.fileIndex.Flush()

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
