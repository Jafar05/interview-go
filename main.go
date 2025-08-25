package main

import "fmt"

func main() {
	// Тестовые случаи
	testCases := []struct {
		nums1, nums2 []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}}, // [2, 2]
	}

	for _, test := range testCases {
		result := intersect(test.nums1, test.nums2)
		fmt.Printf("nums1: %v, nums2: %v -> %v\n", test.nums1, test.nums2, result)
	}
}
func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	res := []int{}
	for _, num := range nums1 {
		hash[num]++
	}
	fmt.Println(hash)
	for _, num := range nums2 {
		if hash[num] > 0 {
			res = append(res, num)
			hash[num]--
		}
	}
	return res
}
