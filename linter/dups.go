package linter

func FindDups(text string) []Token {
	dups := []Token{}
	parser := NewParser(text)
	var previous Token

	for _, token := range parser.Parse() {
		if token.Word == previous.Word {
			dups = append(dups, token)
		}
		previous = token
	}

	return dups
}
