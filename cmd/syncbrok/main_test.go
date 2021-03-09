package main

import (
	"testing"

	"github.com/mes1234/syncbrok/internal/functions"
)

func TestHello(t *testing.T) {
	msg := functions.Format("Gladys")
	if msg != "yo, Gladys. yo!" {
		t.Fatalf("Error in name")
	}
}
