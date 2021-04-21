package msg

import (
	"encoding/json"
	"log"
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

func (m *simpleMsg) RemoveContent() {
	m.Content = nil
}

func (m simpleMsg) GetParentId() uuid.UUID {
	return m.Parent
}

func (m simpleMsg) GetContent() []byte {
	return m.Content
}

func (m simpleMsg) Process(
	wgParent *sync.WaitGroup,
	wgSelf *sync.WaitGroup,
	callback Callback,
	endpoints []string,
	store func(uuid.UUID) []byte,
	storeAck chan<- uuid.UUID) {
	defer wgSelf.Done()
	log.Print("processing begins")
	if m.Parent != uuid.Nil {
		log.Print("I will wait for my parent to finish")
		wgParent.Wait()
		log.Print("my parent completed I shall proceed")
	} else {
		log.Print("I dont have parent let me do my job")
	}
	m.Content = store(m.GetId())
	response, _ := json.Marshal(m)
	status := true
	for _, endpoint := range endpoints {
		status = status && callback(response, endpoint)
	}
	if status {
		storeAck <- m.Id
	} else {
		time.Sleep(3 * time.Second)
		wgSelf.Add(1)
		go m.Process(wgParent, wgSelf, callback, endpoints, store, storeAck)
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
