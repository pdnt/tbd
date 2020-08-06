package linter

// FindDups finds duplicate tokens
func FindDups(parser *Parser) []Token {
	var previous Token
	duplicates := []Token{}

	for _, token := range parser.GetTokens() {
		if token.Kind == SpaceKind {
			continue
		}

		if token.Equals(previous) {
			duplicates = append(duplicates, token)
		}

		previous = token
	}

	return duplicates
}
