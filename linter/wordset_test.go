package linter

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordSet(t *testing.T) {
	var SpaceWords = NewWordSet([]string{"NASA", "Mars", "perseverance"})
	fmt.Println(SpaceWords)
	tests := map[string]struct {
		input string
		want  bool
	}{
		"returns false for words that are not in the set": {input: "Ocean", want: false},
		"true for words that are in the set":              {input: "NASA", want: true},
		"is case insensitive":                             {input: "mars", want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SpaceWords.Has(&Token{Value: tc.input})
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
