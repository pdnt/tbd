package linter

import (
	"regexp"
)

func FindMisspell(parser *Parser) []Token {
	misspellings := []Token{}
	allTokens := parser.GetTokens()
	// FIXME: we should really check if a string is a number in the parser, but we decided to do this quick
	// fix for the time being until we find another use-case that justify the effort
	r, _ := regexp.Compile("^[0-9]+$")

	for _, token := range allTokens {
		if !r.MatchString(token.Value) && token.Kind == WordKind && !DictionaryWords.Has(&token) {
			misspellings = append(misspellings, token)
		}
	}
	return misspellings
}
