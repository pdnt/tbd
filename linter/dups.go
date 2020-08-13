package linter

// FindDups finds duplicate tokens
func FindDups(parser *Parser) []Token {
	var previous Token
	duplicates := []Token{}
	allTokens := parser.GetTokens()

	for i, token := range allTokens {
		if token.Kind == SpaceKind {
			if i > 1 && token.Equals(allTokens[i-1]) {
				duplicates = append(duplicates, token)
			}

			continue
		}

		if token.Equals(previous) {
			duplicates = append(duplicates, token)
		}

		previous = token
	}

	return duplicates
}
