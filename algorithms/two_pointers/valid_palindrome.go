package two_pointers

import (
	"regexp"
	"unicode"
)

func ValidPalindrome() bool {
	s := "A man, a plan, a canal: Panama"

	p1, p2 := 0, len(s)-1

	for p1 <= p2 {
		for p1 <= p2 && !isAlphaNumeric(string(s[p1])) {
			p1++
		}
		for p1 <= p2 && !isAlphaNumeric(string(s[p2])) {
			p2--
		}

		if p1 <= p2 {
			if unicode.ToLower(rune(s[p1])) != unicode.ToLower(rune(s[p2])) {
				return false
			}
			p1++
			p2--
		}

	}

	return true
}

func isAlphaNumeric(w string) bool {
	r := rune(w[0])
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isAlphaNumericReg(w string) bool {
	matched, _ := regexp.MatchString(`^[[:alnum:]]$`, w)
	return matched
}
