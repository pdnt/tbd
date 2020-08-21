package reporters

import (
	"fmt"
	"strings"
	"tbd/linter"
)

// getLineForToken returns the line where the reference token is located.
func getLineForToken(allWords *[]linter.Token, referenceToken *linter.Token) string {
	result := ""

	for _, token := range *allWords {
		if token.Line == referenceToken.Line {
			if token.Kind == linter.SpaceKind && token.Row == referenceToken.Row {
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
	Name    string
	Tokens  []linter.Token
	Enabled bool
}

func ReportStdout(reports []Report, allWords []linter.Token, filePath string) {

	for _, report := range reports {
		if report.Enabled {
			reportPrint(report.Tokens, allWords, report.Name, filePath)
		}
	}
}
