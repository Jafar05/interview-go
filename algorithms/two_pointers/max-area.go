package two_pointers

import "math"

func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	result_area := float64(0)

	for l < r {
		curr_area := math.Min(float64(height[l]), float64(height[r])) * float64((r - l))
		result_area = math.Max(float64(result_area), float64(curr_area))

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return int(result_area)
}
