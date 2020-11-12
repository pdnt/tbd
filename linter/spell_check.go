package linter

import (
	"fmt"
	"regexp"
)

func FindMisspell(parser *Parser, dictionary WordSet) []Token {
	misspellings := []Token{}
	allTokens := parser.GetTokens()
	// FIXME: we should really check if a string is a number in the parser, but we decided to do this quick
	// fix for the time being until we find another use-case that justify the effort
	r, _ := regexp.Compile("^[0-9]+$")

	for _, token := range allTokens {
		fmt.Printf("value: %v | kind: %v\n", token.Value, token.Kind)
		if !r.MatchString(token.Value) && token.Kind == WordKind && !dictionary.Has(&token) {
			misspellings = append(misspellings, token)
		}
	}
	return misspellings
}
