package linter

import (
	"fmt"
	"regexp"
)

func FindDups(text string) []string {
	wordFinderRegex := regexp.MustCompile(`\W+`)
	words := wordFinderRegex.Split(text, -1)
	out := []string{}

	for _, word := range words {
		fmt.Printf("this is a word: %s\n", word)
	}

	return out
}
