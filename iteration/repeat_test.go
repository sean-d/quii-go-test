package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat2(t *testing.T) {
	got := Repeat("b", 0)
	want := "b"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for _ = range b.N {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	//Output: aaaaaaaaaa
}
