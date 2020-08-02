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
			input: "word another, word",
			want: []Token{
				{Value: "word", Line: 1, Row: 1},
				{Value: " ", Line: 1, Row: 5, Kind: SpaceKind},
				{Value: "another", Line: 1, Row: 6},
				{Value: ",", Line: 1, Row: 13, Kind: PunctuationKind},
				{Value: " ", Line: 1, Row: 14, Kind: SpaceKind},
				{Value: "word", Line: 1, Row: 17},
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
