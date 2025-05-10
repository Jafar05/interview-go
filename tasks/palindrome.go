package tasks

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	reverseStr := strings.Builder{}

	for i := len(str) - 1; i >= 0; i-- {
		reverseStr.WriteByte(str[i])
	}

	return str == reverseStr.String()
}

func ImportPalindrome() {
	fmt.Println(isPalindrome("Anna"))
}
