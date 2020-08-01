package linter

// FindDups finds duplicate tokens
func FindDups(parser *Parser) []Token {
	var previous, beforePunctuation Token
	allWords := parser.GetTokens()
	dups := []Token{}

	for _, token := range allWords {
		if token.Value == previous.Value && token.Row != 1 || previous.Kind == PunctuationKind && beforePunctuation.Value == token.Value && token.Kind != SpaceKind {
			dups = append(dups, token)
		}

		if token.Kind == PunctuationKind && previous.Kind != PunctuationKind {
			beforePunctuation = previous
		}

		previous = token
	}

	return dups
}
