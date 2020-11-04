package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"tbd/linter"
	"tbd/reporters"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tbd <glob>",
	Short: "TBD is a document linter, by default it will look for duplicates, weasels and passive voice on the provided files",
	Run: func(cmd *cobra.Command, args []string) {
		for _, filePath := range args {
			// TODO: use a buffer or another better way to read the file contents
			contents, err := ioutil.ReadFile(filePath)

			if err != nil {
				fmt.Printf("there was an error opening the file '%s'\n", filePath)
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if !ReportDuplicates && !ReportWeasel && !ReportPassive && !ReportMisspellings {
				ReportDuplicates = true
				ReportWeasel = true
				ReportPassive = true
				ReportMisspellings = true
			}

			parser := linter.NewParser(string(contents))
			allWords := parser.GetTokens()
			reports := []reporters.Report{
				{Name: "Duplicate", Tokens: linter.FindDups(parser, IncludeWhitespace, IncludePunctuation), Enabled: ReportDuplicates},
				{Name: "Weasel", Tokens: linter.FindWeasel(parser), Enabled: ReportWeasel},
				{Name: "Passive", Tokens: linter.FindPassive(parser), Enabled: ReportPassive},
				{Name: "Misspelling", Tokens: linter.FindMisspell(parser), Enabled: ReportMisspellings},
			}
			reporters.ReportStdout(reports, allWords, filePath)
		}
	},
}

// ReportDuplicates indicates wheter we should report duplicates or not.
var ReportDuplicates bool

// ReportPassive indicates wheter we should report passives or not.
var ReportPassive bool

// ReportWeasel indicates wheter we should report weasels or not.
var ReportWeasel bool

//ReportMisspellings indicates wheter we should report misspellings or not.
var ReportMisspellings bool

// IncludeWhitespace indicates wheter we should include white spaces in duplicates.
var IncludeWhitespace bool

// IncludePunctuation indicates wheter we should consider tokens with punctuation values.
var IncludePunctuation bool

// SetUpFlags sets the CLI usage for the user.
func SetUpFlags() {
	rootCmd.Flags().BoolVarP(&ReportPassive, "passive", "p", false, "Run passives")
	rootCmd.Flags().BoolVarP(&ReportWeasel, "weasel", "w", false, "Run weasels")
	rootCmd.Flags().BoolVarP(&ReportDuplicates, "duplicate", "d", false, "Run duplicates")
	rootCmd.Flags().BoolVarP(&ReportMisspellings, "misspell", "m", false, "Run misspell")
	rootCmd.Flags().BoolVar(&IncludePunctuation, "include-punctuation", false, "Include punctuation")
	rootCmd.Flags().BoolVar(&IncludeWhitespace, "include-whitespace", false, "Include whitespace")
}

// Execute is the entry point for the command line interface
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
