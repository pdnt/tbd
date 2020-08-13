package reporters

import (
	"fmt"
	"strings"
	"tbd/linter"
)

// getLineForToken returns the line where the duplicate word is located.
func getLineForToken(allWords *[]linter.Token, duplicate *linter.Token) string {
	result := ""

	for _, token := range *allWords {
		if token.Line == duplicate.Line {
			if token.Kind == linter.SpaceKind && token.Row == duplicate.Row {
				result += "•"
			} else {
				result += token.Value
			}
		}

	}

	return result
}

// getCarrot returns a carrot ^ below the duplicate word
func getCarrot(duplicate *linter.Token) string {
	return strings.Repeat(" ", duplicate.Row-1) + "^"
}

func reportPrint(tokens, allWords []linter.Token, linter, filePath string) {
	for _, token := range tokens {
		fmt.Printf("%v word in file %s [Line:%v , Row:%v] \n", linter, filePath, token.Line, token.Row)
		fmt.Println(getLineForToken(&allWords, &token))
		fmt.Println(getCarrot(&token))
	}
}

type Report struct {
	Name   string
	Tokens []linter.Token
}

func ReportStdout(parser *linter.Parser, filePath string) {
	allWords := parser.GetTokens()
	reports := []Report{
		{Name: "Duplicate", Tokens: linter.FindDups(parser)},
		{Name: "Weasel", Tokens: linter.FindWeasel(parser)},
	}

	for _, report := range reports {
		reportPrint(report.Tokens, allWords, report.Name, filePath)
	}
}
