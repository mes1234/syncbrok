package msg

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

//Message is an entity passed between channels
type simpleMsg struct {
	id              uuid.UUID
	parent          uuid.UUID
	content         []byte
	deliveryCounter int
}

func (m simpleMsg) GetItems() interface{} {
	return m.content
}

func (m simpleMsg) GetId() uuid.UUID {
	return m.id
}

func (m simpleMsg) GetParentId() uuid.UUID {
	return m.parent
}

func (m simpleMsg) Process(wgParent *sync.WaitGroup, wgSelf *sync.WaitGroup, callbacks []Callback) {
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
	for _, callback := range callbacks {
		callback(m.content)
	}

}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func NewSimpleMsg(parentId uuid.UUID, content []byte) Msg {
	return simpleMsg{
		id:              uuid.New(),
		parent:          parentId,
		content:         content,
		deliveryCounter: 0,
	}
}
