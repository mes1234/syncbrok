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

func TestCleaningContent(t *testing.T) {
	content := []byte{0, 1, 2, 3}
	obj := msg.NewSimpleMsg(uuid.Nil, content)
	obj.RemoveContent()
	if obj.GetContent() != nil {
		t.Errorf("Removing content shall clean content field %v", obj)
	}
}
