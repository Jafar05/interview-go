package two_pointers

func SumTwoTarget(arr []int, target int) []int {

	p1, p2 := 0, len(arr)-1

	for p1 <= p2 {
		sum := arr[p1] + arr[p2]
		if sum == target {
			return []int{p1, p2}
		} else if sum < target {
			p1++
		} else if sum > target {
			p2--
		}
	}

	return []int{-1, -1}
}
