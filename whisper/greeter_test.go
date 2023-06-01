package main

import (
    "testing"
)

func TestGreeting(t *testing.T) {
	if msg := greet(); msg != "" {
		t.Errorf("Unexpected greeting: %s", msg)
	}
}
