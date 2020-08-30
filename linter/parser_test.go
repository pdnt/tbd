package linter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []Token
	}{
		"simple": {
			input: "word another, last",
			want: []Token{
				{Value: "word", Line: 1, Row: 1, Index: -1},
				{Value: " ", Line: 1, Row: 5, Kind: SpaceKind, Index: 0},
				{Value: "another", Line: 1, Row: 6, Index: 1},
				{Value: ",", Line: 1, Row: 13, Kind: PunctuationKind, Index: 2},
				{Value: " ", Line: 1, Row: 14, Kind: SpaceKind, Index: 3},
				{Value: "last", Line: 1, Row: 15, Index: 4},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(tc.input)
			got := parser.GetTokens()
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
