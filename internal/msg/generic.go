package msg

import (
	"sync"

	"github.com/google/uuid"
)

type Callback func([]byte, string) bool

//Msg is main entinty
type Msg interface {
	GetWaiter() *sync.WaitGroup
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, Callback, []string, chan<- uuid.UUID)
	GetContent() []byte
	GetEpochs() int64
}
