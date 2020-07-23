package reporters

import (
	"fmt"
	"tbd/linter"
)

func ReportCl(parser *linter.Parser) {
	duplicates := linter.FindDups(parser)

	// Add context
	fmt.Println(duplicates)

}
