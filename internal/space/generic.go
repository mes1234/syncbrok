package space

import (
	"github.com/mes1234/syncbrok/internal/queue"
)

type Space interface {
	AddQueue(queue.Queue, string)
}
