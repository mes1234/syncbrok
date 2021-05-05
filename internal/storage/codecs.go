package storage

import (
	"bytes"
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
