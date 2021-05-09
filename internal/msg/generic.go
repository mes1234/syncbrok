package msg

import (
	"sync"

	"github.com/google/uuid"
)

type CallbackFactory func(func(uuid.UUID) []byte) Callback
type Callback func(uuid.UUID, string, *sync.WaitGroup)

//Msg is main entinty
type Msg interface {
	GetWaiter() *sync.WaitGroup
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, Callback, []string)
	GetContent() []byte
	RemovePayload()
}

type MsgContent struct {
	Id      uuid.UUID
	Content []byte
}
