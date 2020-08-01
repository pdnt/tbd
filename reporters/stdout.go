package reporters

import (
	"fmt"
	"strings"
	"tbd/linter"
)

func getLineForToken(allWords *[]linter.Token, duplicate *linter.Token) string {
	result := ""

	for _, token := range *allWords {
		if token.Line == duplicate.Line {
			result += token.Value
		}
	}

	return result
}

func getCarrot(duplicate *linter.Token) string {
	return strings.Repeat(" ", duplicate.Row-1) + "^"
}

func ReportStdout(parser *linter.Parser, filePath string) {
	allWords := parser.GetTokens()
	duplicates := linter.FindDups(parser)

	for _, duplicate := range duplicates {
		fmt.Printf("Duplicated word in file %s [Line:%v , Row:%v] \n", filePath, duplicate.Line, duplicate.Row)
		fmt.Println(getLineForToken(&allWords, &duplicate))
		fmt.Println(getCarrot(&duplicate))
	}
}
