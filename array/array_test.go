package array

import (
	"testing"
)

func TestArray(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Array2(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
