package storage

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

func (fw *FileWriter) addToStore(m msg.Msg) {
	content := m.GetContent()
	msg := MsgSave{
		StartOffset: fw.offset,
		Len:         len(content),
		Id:          m.GetId(),
		Parent:      m.GetParentId(),
	}
	fw.lookup[m.GetId()] = msg

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
