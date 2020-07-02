package linter

import (
	"unicode"
)

type Token struct {
	Word string
	Line int
	Row  int
}

type Parser struct {
	text   string
	tokens []Token
}

func NewParser(text string) *Parser {
	parser := &Parser{
		text: text,
	}

	return parser
}

func (p *Parser) AddToken(word string, line, row int) {
	token := &Token{
		Word: word,
		Line: line,
		Row:  row - len(word),
	}

	p.tokens = append(p.tokens, *token)
}

func (p *Parser) Parse() []Token {
	line := 1
	word := ""
	row := 0

	for _, code := range p.text {
		// If newline, increment line and resets row value. Else increments row value.
		row++
		if code == 10 {
			p.AddToken(word, line, row)
			line++
			row = 0
			word = ""
			continue
		}

		// If new letter, add letter to word.
		if unicode.IsLetter(code) || unicode.IsSymbol(code) || unicode.IsPunct(code) || unicode.IsNumber(code) {
			word = word + string(code)
			continue
		}

		// If spacebreak, create token and allocate the word.
		if unicode.IsSpace(code) {
			p.AddToken(word, line, row)
			word = ""
			continue
		}

	}

	return p.tokens
}
