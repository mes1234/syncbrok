package msg

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

//Message is an entity passed between channels
type simpleMsg struct {
	id      uuid.UUID
	parent  uuid.UUID
	content []byte
}

func (m simpleMsg) GetId() uuid.UUID {
	return m.id
}

func (m *simpleMsg) RemoveContent() {
	m.content = nil
}

func (m simpleMsg) GetParentId() uuid.UUID {
	return m.parent
}

func (m simpleMsg) GetContent() []byte {
	return m.content
}

func (m simpleMsg) Process(
	wgParent *sync.WaitGroup,
	wgSelf *sync.WaitGroup,
	callback Callback,
	endpoints []string,
	store func(uuid.UUID) []byte) {
	defer wgSelf.Done()
	log.Print("processing begins")
	if m.parent != uuid.Nil {
		log.Print("I will wait for my parent to finish")
		wgParent.Wait()
		log.Print("my parent completed I shall proceed")
	} else {
		log.Print("I dont have parent let me do my job")
		time.Sleep(20 * time.Second)
	}
	m.content = store(m.GetId())
	for _, endpoint := range endpoints {
		callback(m.content, endpoint)
	}

}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func NewSimpleMsg(parentId uuid.UUID, content []byte) Msg {
	return &simpleMsg{
		id:      uuid.New(),
		parent:  parentId,
		content: content,
	}
}
