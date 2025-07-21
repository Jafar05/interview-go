package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 22))
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for idx, val := range nums {
		diff := target - val
		if findNum, ok := hash[diff]; ok {
			return []int{findNum, idx}
		}

		hash[val] = idx
	}
	return []int{-1, -1}
}
