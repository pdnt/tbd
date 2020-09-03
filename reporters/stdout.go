package reporters

import (
	"fmt"
	"strings"
	"tbd/linter"
)

func reduceToLine(allWords *[]linter.Token, referenceToken *linter.Token) *[]linter.Token {
	ctx := 6
	arr := []linter.Token{}
	lowerBoundry := referenceToken.Index - ctx
	upperBoundry := referenceToken.Index + ctx

	if lowerBoundry < 0 {
		lowerBoundry = 0
	}

	if upperBoundry > len(*allWords)-1 {
		upperBoundry = len(*allWords) - 1
	}

	for i := lowerBoundry; i <= upperBoundry; i++ {
		token := (*allWords)[i]
		if token.Line == referenceToken.Line {
			arr = append(arr, token)
		}
	}

	return &arr
}

// getLineForToken returns the line where the reference token is located.
func getLineForToken(surroundingWords *[]linter.Token, referenceToken *linter.Token) string {
	result := ""

	for _, token := range *surroundingWords {
		if token.Line == referenceToken.Line {
			if token.Kind == linter.SpaceKind && token.Row == referenceToken.Row {
				result += "•"
			} else if token.DeepEquals(*referenceToken) {
				var boldValue = Colorize(Bold, token.Value)
				result += Colorize(Red, boldValue)
			} else {
				result += token.Value
			}
		}
	}

	return result
}

// getCarrot returns a carrot ^ below the duplicate word
func getCarrot(surroundingWords *[]linter.Token, referenceToken *linter.Token) string {
	carrot := ""
	for _, token := range *surroundingWords {
		if token.DeepEquals(*referenceToken) {
			break
		}

		carrot += strings.Repeat(" ", len(token.Value))
	}
	carrot = Colorize(Bold, carrot+"^")
	return Colorize(Red, carrot)
}

func reportPrint(tokens, allWords []linter.Token, linter, filePath string) {
	fmt.Printf(Colorize(Reverse, fmt.Sprintf("            %vs            \n", linter)) + "\n")
	for _, token := range tokens {
		surroundingWords := reduceToLine(&allWords, &token)
		fmt.Printf("File %s [line:%v , col:%v] \n", filePath, token.Line, token.Row)
		fmt.Println(getLineForToken(surroundingWords, &token))
		fmt.Println(getCarrot(surroundingWords, &token))
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
