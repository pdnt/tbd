package linter

import "testing"

func TestFindDups(t *testing.T) {
	str := "This is a string. It contains dups dups, commas, Upper upper dups"
	out := FindDups(str)

	if len(out) != 2 {
		t.Errorf("expected to get two dups, but got %d", len(out))
	}
}
