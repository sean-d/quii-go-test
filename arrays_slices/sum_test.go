package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of a number of ints", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("got %q want %q from %#v", got, want, nums)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("summing up multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{1, 2, 3, 4, 5})
		want := []int{3, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("summing all tails with greater than 2 items", func(t *testing.T) {
		got := SumAllTails([]int{0, 1, 2}, []int{1, 4})
		want := []int{3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("summing all tails with less than 2 items", func(t *testing.T) {
		got := SumAllTails([]int{2}, []int{1, 4})
		want := []int{2, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("summing all tails with 0 items", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 4})
		want := []int{0, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
