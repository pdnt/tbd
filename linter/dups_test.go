package linter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindDups(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []Token
	}{
		"simple":      {input: "this contains a dup dup word", want: []Token{{Value: "dup", Line: 1, Row: 21, Index: 7}}},
		"punctuation": {input: "this contains,, duplicate punct", want: []Token{{Value: ",", Line: 1, Row: 15, Kind: PunctuationKind, Index: 3}}},
		"spaces":      {input: "this contains  duplicate spaces", want: []Token{{Value: " ", Line: 1, Row: 15, Kind: SpaceKind, Index: 3}}},
		"newline":     {input: "words after newline\nnewline are not dups", want: []Token{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(tc.input)
			got := FindDups(parser)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
