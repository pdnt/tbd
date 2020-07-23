package linter

import (
	"unicode"
)

const (
	WordKind = iota
	PunctuationKind
)

type Token struct {
	Value string
	Line  int
	Row   int
	Kind  int
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

func (p *Parser) AddToken(value string, line, row, kind int) {
	token := &Token{
		Value: value,
		Line:  line,
		Row:   row - len(value),
		Kind:  kind,
	}

	p.tokens = append(p.tokens, *token)
}

func (p *Parser) Parse() []Token {
	value := ""
	line := 1
	row := 0

	for _, code := range p.text {
		// If newline, increment line and resets row value. Else increments row value.
		row++
		if code == 10 {
			p.AddToken(value, line, row, WordKind)
			line++
			row = 0
			value = ""
			continue
		}

		if unicode.IsPunct(code) {
			p.AddToken(string(code), line, row, PunctuationKind)
			continue
		}

		// If new letter, add letter to word.
		if unicode.IsLetter(code) || unicode.IsSymbol(code) || unicode.IsNumber(code) {
			value = value + string(code)
			continue
		}

		// If spacebreak, create token and allocate the word.
		if unicode.IsSpace(code) {
			p.AddToken(value, line, row, WordKind)
			value = ""
			continue
		}

	}

	return p.tokens
}
