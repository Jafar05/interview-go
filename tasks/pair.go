package tasks

func isPair(s string) bool {
	stask := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stask = append(stask, char)
		case ')', '}', ']':
			if len(stask) != 0 && stask[len(stask)-1] != pairs[char] {
				return false
			}
			stask = stask[:len(stask)-1]
		}
	}
	return len(stask) == 0
}
