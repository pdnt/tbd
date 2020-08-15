package linter

// FindWeasel find weasel tokens
func FindWeasel(parser *Parser) []Token {
	weasels := []Token{}
	allTokens := parser.GetTokens()
	allTokensLen := len(allTokens)

	for i, token := range allTokens {
		if WeaselWords.Has(&token) {
			weasels = append(weasels, token)
			continue
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
