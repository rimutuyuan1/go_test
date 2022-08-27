package iteration

import "testing"

func TestInteration(t *testing.T) {
	repeated := Iteration("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
