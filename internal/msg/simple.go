package msg

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
)

//Message is an entity passed between channels
type simpleMsg struct {
	Id      uuid.UUID `json:"id"`
	Parent  uuid.UUID `json:"parent"`
	Content []byte    `json:"content"`
}

func (m simpleMsg) GetId() uuid.UUID {
	return m.Id
}

func (m simpleMsg) GetParentId() uuid.UUID {
	return m.Parent
}

func (m simpleMsg) GetContent() []byte {
	return m.Content
}

func (m simpleMsg) Process(wgParent *sync.WaitGroup, wgSelf *sync.WaitGroup, callback Callback, endpoints []string, ack chan<- uuid.UUID) {
	defer wgSelf.Done()
	if m.Parent != uuid.Nil {
		wgParent.Wait()
	}
	content, _ := json.Marshal(m)
	status := true
	for _, endpoint := range endpoints {
		status = status && callback(content, endpoint)
	}
	if status {
		ack <- m.Id
	} else {
		time.Sleep(3 * time.Second)
		wgSelf.Add(1)
		go m.Process(wgParent, wgSelf, callback, endpoints, ack)
	}

}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func NewSimpleMsg(parentId uuid.UUID, content []byte) Msg {
	return &simpleMsg{
		Id:      uuid.New(),
		Parent:  parentId,
		Content: content,
	}
}
