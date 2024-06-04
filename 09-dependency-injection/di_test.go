package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tiago")

	got := buffer.String()
	want := "Hello, Tiago"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
