package msg

import (
	"sync"

	"github.com/google/uuid"
)

type Callback func([]byte, string) bool

//Msg is main entinty
type Msg interface {
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, *sync.WaitGroup, Callback, []string, chan<- uuid.UUID)
	GetContent() []byte
}
