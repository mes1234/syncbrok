package msg

import "github.com/google/uuid"

//Message is an entity passed between channels
type Msg struct {
	Id              uuid.UUID
	Parent          *Msg
	Content         interface{}
	DeliveryCounter int
}

//Init initilizes new message for given parent and content
//automatically assing global uniq uuid
func New(parent *Msg, content interface{}) Msg {
	return Msg{
		Id:              uuid.New(),
		Parent:          parent,
		Content:         content,
		DeliveryCounter: 0}
}
