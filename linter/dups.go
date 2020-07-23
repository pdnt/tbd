package linter

func FindDups(parser *Parser) []Token {
	dups := []Token{}
	var previous Token
	allWords := parser.Parse()

	for _, token := range allWords {
		if token.Value == previous.Value {
			dups = append(dups, token)
		}
		previous = token

	}

	return dups
}
