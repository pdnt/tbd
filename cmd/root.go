package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"tbd/linter"

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
				fmt.Printf("there was an error opening %s", filePath)
				os.Exit(1)
			}

			linter.FindDups(string(contents))
		}
	},
}

// Execute is the entry point for the command line interface
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}