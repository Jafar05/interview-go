package Задачи

func BinarySearch(nums []int, target int) (int, bool) {
	first, last := 0, len(nums)-1

	for first <= last {
		mid := ((last - first) / 2) + first

		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < target {
			first = mid + 1
		} else if nums[mid] > target {
			last = mid - 1
		}
	}
	return -1, false
}
