package msg

import "github.com/google/uuid"

type IMsg interface {
	GetItems() interface{}
	GetId() uuid.UUID
}

//Message is an entity passed between channels
type Msg struct {
	Id              uuid.UUID
	Parent          *IMsg
	Content         interface{}
	DeliveryCounter int
}

func (m Msg) GetItems() interface{} {
	return m.Content
}

func (m Msg) GetId() uuid.UUID {
	return m.Id
}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func New(parent *IMsg, content interface{}) IMsg {
	return Msg{
		Id:              uuid.New(),
		Parent:          parent,
		Content:         content,
		DeliveryCounter: 0}
}
