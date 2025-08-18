package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(countWord("Go is awesome! 123, but why?"))
}

func countWord(text string) map[string]int {
	hash := make(map[string]int)

	text = strings.ToLower(text)

	reg := regexp.MustCompile(`[^a-zа-я\s]`)
	text = reg.ReplaceAllString(text, "")

	newString := strings.Fields(text)

	for _, char := range newString {
		if _, ok := hash[char]; !ok {
			hash[char] = 1
		}
		hash[char]++
	}
	return hash
}
