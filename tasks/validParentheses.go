package tasks

func validParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) != 0 && stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
