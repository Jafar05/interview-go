package slice

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func MainUniqN() {
	fmt.Println(uniqN(10))
}

func uniqN(n int) []int {
	sl := make([]int, 0, n)
	uniqMap := make(map[int]struct{})
	for len(sl) < n {
		bi, _ := rand.Int(rand.Reader, big.NewInt(15))
		r := int(bi.Int64())
		if _, ok := uniqMap[r]; !ok {
			sl = append(sl, r)
			uniqMap[r] = struct{}{}
		}
	}
	fmt.Println(len(sl))
	return sl
}
