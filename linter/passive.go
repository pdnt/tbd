package linter

import (
	"strings"
)

func nextWord(tokens *[]Token, startRange int) *Token {
	for i := startRange; i < len(*tokens); i++ {
		if (*tokens)[i].Kind == WordKind {
			return &(*tokens)[i]
		}
	}

	return nil
}

// FindPassive finds tokens using pasive voice
func FindPassive(parser *Parser) []Token {
	var passive []Token
	allTokens := parser.GetTokens()

	for i, token := range allTokens {
		if token.Kind == WordKind && LinkingVerbs.Has(&token) {
			nw := nextWord(&allTokens, i+1)
			if strings.HasSuffix(nw.Value, "ed") || IrregularWords.Has(&token) {
				passive = append(passive, token)
			}
		}
	}

	return passive
}
