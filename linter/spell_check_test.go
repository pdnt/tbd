package linter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSpellCheck(t *testing.T) {
	dictionaryWordsMock := NewWordSet([]string{"is", "a", "misspelling", "number", "on", "the", "test", "there", "are", "no", "errors", ""})

	tests := map[string]struct {
		input string
		want  []Token
	}{
		"misspelling": {input: "tis is a misspelling", want: []Token{{Value: "tis", Line: 1, Row: 1, Index: -1}}},
		"is number":   {input: "there is a 1 number 200 on the test!", want: []Token{}},
		"no errors":   {input: "There are no errors on the test.", want: []Token{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parser := NewParser(tc.input)
			got := FindMisspell(parser, dictionaryWordsMock)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
