package msg

import (
	"sync"

	"github.com/google/uuid"
)

type Callback func([]byte, string) bool

type Msg interface {
	GetItems() interface{}
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, *sync.WaitGroup, Callback, []string)
	GetContent() []byte
}
