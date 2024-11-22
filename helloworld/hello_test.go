package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("to you")
	want := "Hello, to you"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
