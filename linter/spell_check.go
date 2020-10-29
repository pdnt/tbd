package linter

func FindMisspell(parser *Parser) []Token {
	misspellings := []Token{}
	allTokens := parser.GetTokens()

	for _, token := range allTokens {
		if token.Kind == WordKind && !DictionaryWords.Has(&token) {
			misspellings = append(misspellings, token)
		}
	}
	return misspellings
}
