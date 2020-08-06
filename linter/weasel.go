package linter

import (
	"strings"
)

func FindWeasel(parser *Parser) []Token {
	allWeasels := map[string]bool{"many": true, "various": true, "very": true, "fairly": true, "several": true, "extremely": true, "exceedingly": true, "quite": true, "remarkably": true, "few": true, "surprisingly": true, "mostly": true, "largely": true, "huge": true, "tiny": true /*"((are","is)"a"number)"*/, "excellent": true, "interestingly": true, "significantly": true, "substantially": true, "clearly": true, "vast": true, "relatively": true, "completely": true}
	weasels := []Token{}
	allTokens := parser.GetTokens()
	allTokensLen := len(allTokens)

	for i, token := range allTokens {
		isInWeasels, _ := allWeasels[strings.ToLower(token.Value)]

		if isInWeasels {
			weasels = append(weasels, token)
		}

		if (token.ValueEquals("is") || token.ValueEquals("are")) && i+4 < allTokensLen {
			w1 := allTokens[i+2]
			w2 := allTokens[i+4]

			if w1.Equals(Token{Value: "a", Line: token.Line}) && w2.Equals(Token{Value: "number", Line: token.Line}) {
				weasels = append(weasels, token)
			}
		}
	}
	return weasels
}
