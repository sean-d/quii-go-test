package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
