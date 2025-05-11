package two_pointers

import "fmt"

func MoveZeroes( /* []int{0, 1, 0, 0, 3, 4, 0} */ ) {
	nums := []int{0, 1, 0, 0, 3, 4, 0}
	freeIdx := 0

	for _, num := range nums {
		if num != 0 {
			nums[freeIdx] = num
			freeIdx++
		}
	}

	for i := freeIdx; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}
