package storage

import (
	"encoding/gob"
	"log"
)

func (fw *FileWriter) encodeMsgSave(msg MsgSave) (bytes []byte) {
	encoder := gob.NewEncoder(&fw.buffer)
	err := encoder.Encode(msg)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	bytes = fw.buffer.Bytes()
	fw.buffer.Reset()
	return
}
