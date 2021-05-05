package msg_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mes1234/syncbrok/internal/msg"
)

func TestNewSimpleMsgCreation(t *testing.T) {
	got := msg.NewSimpleMsg(uuid.Nil, nil)
	if got.GetId() == uuid.Nil {
		t.Errorf("Gid shall be assigned for evey msg %v", got)
	}
}
