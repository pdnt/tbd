package linter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPassive(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []Token
	}{
		"verb + irregular":         {input: "I noticed that a window had been left open.", want: []Token{{Value: "been", Line: 1, Row: 29, Index: 11}}},
		"verb + word ending in ed": {input: "every year thousands of people are killed on our roads.", want: []Token{{Value: "are", Line: 1, Row: 32, Index: 9}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(tc.input)
			got := FindPassive(parser)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
