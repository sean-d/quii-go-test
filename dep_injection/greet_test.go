package dep_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "you")

	got := buffer.String()
	want := "hello, you"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
