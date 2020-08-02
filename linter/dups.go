package linter

// FindDups finds duplicate tokens
func FindDups(parser *Parser) []Token {
	var previous Token
	var duplicates []Token

	for _, token := range parser.GetTokens() {
		if token.Kind == SpaceKind {
			continue
		}

		if previous.Value == token.Value && token.Row != 1 {
			duplicates = append(duplicates, token)
		}

		previous = token
	}

	return duplicates
}
