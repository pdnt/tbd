package linter

import (
	"strings"
	"unicode"
)

const (
	WordKind = iota
	PunctuationKind
	SpaceKind
)

type Token struct {
	Value string
	Line  int
	Row   int
	Kind  int
	Index int
}

func (t *Token) DeepEquals(token Token) bool {
	return t.Equals(token) && t.Row == token.Row
}

// Equals compares two tokens and returns `true` if they are semantically equal.
func (t *Token) Equals(token Token) bool {
	return t.ValueEquals(token.Value) && t.Line == token.Line
}

// ValueEquals compares a given value with the Value property of the token and returns `true` if they are equal.
func (t *Token) ValueEquals(value string) bool {
	return strings.ToLower(t.Value) == strings.ToLower(value)
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
		Index: len(p.tokens) - 1,
	}

	p.tokens = append(p.tokens, *token)
}

func (p *Parser) GetTokens() []Token {
	if len(p.tokens) == 0 {
		return p.parse()
	}

	return p.tokens
}

func (p *Parser) parse() []Token {
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
			if value != "" {
				// First, add the word
				p.AddToken(value, line, row, WordKind)
			}

			// Then add the punctuation sign
			p.AddToken(string(code), line, row+1, PunctuationKind)
			value = ""
			continue
		}

		// If new letter, add letter to word.
		if unicode.IsLetter(code) || unicode.IsSymbol(code) || unicode.IsNumber(code) {
			value = value + string(code)
			continue
		}

		// If spacebreak, create token and allocate the word.
		if unicode.IsSpace(code) {
			if value != "" {
				p.AddToken(value, line, row, WordKind)
			}

			p.AddToken(string(code), line, row+1, SpaceKind)
			value = ""
			continue
		}
	}

	p.AddToken(value, line, row+1, WordKind)

	return p.tokens
}
