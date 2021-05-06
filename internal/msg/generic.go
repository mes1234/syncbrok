package msg

import (
	"sync"

	"github.com/google/uuid"
)

type Callback func([]byte, string, *sync.WaitGroup)

//Msg is main entinty
type Msg interface {
	GetWaiter() *sync.WaitGroup
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, Callback, []string)
	GetContent() []byte
}
