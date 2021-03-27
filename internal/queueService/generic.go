package queueService

import "github.com/mes1234/syncbrok/internal/space"

type QueueService interface {
	Start(space.Space)
}
