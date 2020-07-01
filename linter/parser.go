package linter

import (
	"unicode"
)

type Token struct {
	Word string
	Line int
	Row  int
}

// TODO
//Fix line and word output values

func Parse(text string) []Token {
	tokens := []Token{}
	line := 0
	word := ""
	row := 0

	for _, code := range text {
		// If newline, increment line and resets row value. Else increments row value.
		if code == 10 {
			line++
			row = 0
		}
		// If new letter, add letter to word.
		if unicode.IsLetter(code) || unicode.IsSymbol(code) || unicode.IsPunct(code) {
			word = word + string(code)
		}
		// If spacebreak, create token and allocate the word.
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
