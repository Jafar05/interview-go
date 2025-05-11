package two_pointers

import (
	"fmt"
	"math"
	"sort"
)

func Squaring() {
	nums := []int{-4, -1, 0, 3, 10}
	var res []int
	p1, p2 := 0, len(nums)-1

	for p1 <= p2 {
		if nums[p1] > nums[p2] {
			res = append(res, int(math.Pow(float64(nums[p1]), 2)))
			p1++
		} else {
			res = append(res, int(math.Pow(float64(nums[p2]), 2)))
			p2--
		}
	}

	sort.Ints(res)
	fmt.Println(res)
}
