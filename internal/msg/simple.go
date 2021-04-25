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
	Id        uuid.UUID `json:"id"`
	Parent    uuid.UUID `json:"parent"`
	Content   []byte    `json:"content"`
	Waiter    *sync.WaitGroup
	TimeStamp time.Time `json:"timestamp"`
	Delivered bool
}

func (m *simpleMsg) GetTime() time.Time {
	return m.TimeStamp
}

func (m *simpleMsg) GetWaiter() *sync.WaitGroup {
	return m.Waiter
}

func (m *simpleMsg) GetId() uuid.UUID {
	return m.Id
}

func (m *simpleMsg) GetParentId() uuid.UUID {
	return m.Parent
}

func (m *simpleMsg) GetContent() []byte {
	return m.Content
}

func (m *simpleMsg) Process(wgParent *sync.WaitGroup, callback Callback, endpoints []string, ack chan<- uuid.UUID) {
	defer m.Waiter.Done()
	if wgParent != nil {
		log.Printf("I have parent so I need to wait %v", m.Id)
		wgParent.Wait()
		log.Printf("My parent finished let me proceed %v", m.Id)
	}
	log.Printf("Processing %v", m.Id)
	content, _ := json.Marshal(m)
	status := true
	for _, endpoint := range endpoints {
		status = status && callback(content, endpoint)
	}
	if status {
		ack <- m.Id
		m.Delivered = true
		log.Printf("Finished %v", m.Id)
	} else {
		log.Printf("Not reciving ack try once again %v", m.Id)
		time.Sleep(3 * time.Second)
		m.Waiter.Add(1)
		go m.Process(wgParent, callback, endpoints, ack)
	}

}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func NewSimpleMsg(parentId uuid.UUID, content []byte) Msg {
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	return &simpleMsg{
		Id:        uuid.New(),
		Parent:    parentId,
		Content:   content,
		Waiter:    &waiter,
		TimeStamp: time.Now(),
	}
}
