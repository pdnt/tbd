package linter

import "fmt"

func FindDups(text string) []string {
	// wordFinderRegex := regexp.MustCompile(`\W+`)
	// words := wordFinderRegex.Split(text, -1)
	out := []string{}
	fmt.Println(Parse(text))
	// for _, word := range words {
	// 	fmt.Printf("this is a word: %s\n", word)
	// }

	return out
}
