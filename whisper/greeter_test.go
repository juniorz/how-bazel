package main

import (
    "testing"
	"strings"
)

func TestGreeting(t *testing.T) {
	if msg := greet(); !strings.Contains(msg, "Hello, world from"){
		t.Errorf("Unexpected greeting: %s", msg)
	}
}
