package main

import (
	"testing"
)

// TestHello tests that saying hello to someone (supplied string) or to nobody (empty string supplied) works.
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("english", "to you")
		want := "Hello, to you"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("spanish", "Pedro")
		want := "Hola, Pedro"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("french", "Jean Luc")
		want := "Bonjour, Jean Luc"

		assertCorrectMessage(t, got, want)
	})

}

// assertCorrectMessage takes in t from the Test functions, as well as the got/want, and does the asserting
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
