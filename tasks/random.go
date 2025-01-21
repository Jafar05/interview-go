package tasks

import (
	"math/rand"
)

func uniqRandn(n int) []int {
	sl := make([]int, 0, n)
	m := make(map[int]struct{})
	for len(sl) < n {
		s := rand.Int()
		if _, ok := m[s] {
			continue
		}

		sl = append(sl, s)
		m[s] = struct{}{}
	}

	return sl
}
