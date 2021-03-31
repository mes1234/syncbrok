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
	parent          Msg
	content         interface{}
	deliveryCounter int
	Callback        func(uuid.UUID)
}

func (m simpleMsg) GetItems() interface{} {
	return m.content
}

func (m simpleMsg) GetId() uuid.UUID {
	return m.id
}

func (m simpleMsg) GetParentId() uuid.UUID {
	if m.parent != nil {
		return m.parent.GetId()
	} else {
		return uuid.Nil
	}

}

func (m simpleMsg) Process(wgParent *sync.WaitGroup, wgSelf *sync.WaitGroup) {
	defer wgSelf.Done()
	log.Print("processing begins")
	if m.parent != nil {
		log.Print("I will wait for my parent to finish")
		wgParent.Wait()
		log.Print("my parent completed I shall proceed")
	} else {
		log.Print("I dont have parent let me do my job")
		time.Sleep(5 * time.Second)
	}

	m.Callback(m.id)
}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func NewSimpleMsg(parent Msg, content interface{}, handler func(uuid.UUID)) Msg {
	return simpleMsg{
		id:              uuid.New(),
		parent:          parent,
		content:         content,
		deliveryCounter: 0,
		Callback:        handler,
	}
}
