package linter

import "strings"

func FindWeasel(parser *Parser) []Token {
	allWeasels := []string{"many", "various", "very", "fairly", "several", "extremely", "exceedingly", "quite", "remarkably", "few", "surprisingly", "mostly", "largely", "huge", "tiny" /*"((are","is)"a"number)"*/, "excellent", "interestingly", "significantly", "substantially", "clearly", "vast", "relatively", "completely"}
	weasels := []Token{}
	allWords := parser.GetTokens()

	for _, token := range allWords {
		for _, weasel := range allWeasels {
			if strings.ToLower(token.Value) == weasel {
				weasels = append(weasels, token)

			if  == "is" || item == "are" {
				slice := allWords[]
			}

			}
		}
	}
	return weasels
}
