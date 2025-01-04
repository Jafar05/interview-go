package tasks

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	if len(s1) != len(s2) {
		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	sort.Slice(r1, func(a, b int) bool { return r1[a] < r1[b] })
	sort.Slice(r2, func(a, b int) bool { return r2[a] < r2[b] })

	return string(r1) == string(r2)
}

func ImportAnagrams() {
	fmt.Println(isAnagram("липа", "пила"))
}
