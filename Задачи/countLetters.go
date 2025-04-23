package Задачи

import (
	"fmt"
	"strings"
)

func countLetters(input string) string {
	if len(input) == 0 {
		return ""
	}

	prev := input[0]
	var result strings.Builder
	count := 1

	for i := 0; i < len(input); i++ {
		if input[i] == prev {
			count++
		} else {
			result.WriteString(fmt.Sprintf("%c%d", prev, count))
			prev = input[i]
			count = 1
		}
	}
	result.WriteString(fmt.Sprintf("%c%d", prev, count))
	return result.String()
}

func ImportCountLetters() {
	input := "AAABBAABA"
	result := countLetters(input)
	fmt.Println(result)
}
