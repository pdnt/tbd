package linter

// FindDups finds duplicate tokens
func FindDups(parser *Parser, includeWhitespace, includePunctuation bool) []Token {
	var lastParsedWord Token
	duplicates := []Token{}
	allTokens := parser.GetTokens()

	for i, token := range allTokens {
		if (token.Kind == SpaceKind && includeWhitespace) || (token.Kind == PunctuationKind && includePunctuation) {
			if i > 1 && token.Equals(allTokens[i-1]) {
				duplicates = append(duplicates, token)
			}
		}

		if token.Kind == WordKind {
			if token.Equals(lastParsedWord) {
				duplicates = append(duplicates, token)
			}
			lastParsedWord = token
		}
	}

	return duplicates
}
