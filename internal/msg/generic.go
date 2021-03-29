package msg

import "github.com/google/uuid"

type Msg interface {
	GetItems() interface{}
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(<-chan bool, chan<- bool)
}
