package linter

import "testing"

func TestFindDups(t *testing.T) {
	str := "This is a string. It contains dups dups, commas, Upper upper dups"
	out := FindDups(str)

	if len(out) != 2 {
		t.Errorf("expected to get two dups, but got %d", len(out))

	}
	want := "dups"

	if out[0] != want {
		t.Errorf("Expected the first item to be %v, but got %v", want, out[0])
	}
	want = "upper"
	if out[1] != want {
		t.Errorf("Expected the second item to be %v, but got %v", want, out[1])
	}

}
