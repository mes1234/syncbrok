package msg

import "github.com/google/uuid"

type IMsg interface {
	GetItems() interface{}
	GetId() uuid.UUID
}

//Message is an entity passed between channels
type simpleMsg struct {
	id              uuid.UUID
	parent          *IMsg
	content         interface{}
	deliveryCounter int
}

func (m simpleMsg) GetItems() interface{} {
	return m.content
}

func (m simpleMsg) GetId() uuid.UUID {
	return m.id
}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func New(parent *IMsg, content interface{}) IMsg {
	return simpleMsg{
		id:              uuid.New(),
		parent:          parent,
		content:         content,
		deliveryCounter: 0}
}
