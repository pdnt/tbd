package linter

import (
	"unicode"
)

type Token struct {
	Word string
	Line int
	Row  int
}

func Parse(text string) []Token {
	tokens := []Token{}
	line := 0
	word := ""
	row := 0

	for _, code := range text {
		// if newline, increment line
		if code == 10 {
			line++
			row = 0
		}

		if unicode.IsLetter(code) || unicode.IsSymbol(code) || unicode.IsPunct(code) {
			word = word + string(code)
		}

		if unicode.IsSpace(code) {
			token := &Token{
				Word: word,
				Line: line,
				Row:  row,
			}

			tokens = append(tokens, *token)
			word = ""
		}

		row++
	}

	return tokens
}
