package msg

import (
	"sync"

	"github.com/google/uuid"
)

type Msg interface {
	GetItems() interface{}
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	Process(*sync.WaitGroup, *sync.WaitGroup)
}
