package linter

import (
	"fmt"
)

// FindDups finds duplicate tokens
func FindDups(parser *Parser) []Token {
	var previous, beforePunctuation Token
	allWords := parser.Parse()
	dups := []Token{}

	for _, token := range allWords {
		fmt.Println(token.Value)
		if token.Value == previous.Value && token.Row != 1 || previous.Kind == PunctuationKind && beforePunctuation.Value == token.Value {
			dups = append(dups, token)
		}

		if token.Kind == PunctuationKind && previous.Kind != PunctuationKind {
			beforePunctuation = previous
		}

		previous = token
	}

	return dups
}
