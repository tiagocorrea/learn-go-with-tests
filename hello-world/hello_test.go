package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tiago")
	want := "Hello, Tiago"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
