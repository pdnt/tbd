package linter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWeasel(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []Token
	}{
		"simple":     {input: "clearly this contains a weasel", want: []Token{{Value: "clearly", Line: 1, Row: 1}}},
		"are number": {input: "there are a number of tests", want: []Token{{Value: "are", Line: 1, Row: 7}}},
		"is number":  {input: "There is a number of test", want: []Token{{Value: "is", Line: 1, Row: 7}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(tc.input)
			got := FindWeasel(parser)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
