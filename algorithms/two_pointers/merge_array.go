package two_pointers

func MergeArray() []int {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 10}
	arr2 := []int{1, 5, 6, 7, 10}

	p1, p2 := 0, 0

	var res []int

	for p1 < len(arr1) && p2 < len(arr2) {
		if arr1[p1] == arr2[p2] {
			res = append(res, arr1[p1])
			p1++
			p2++
		} else if arr1[p1] < arr2[p2] {
			p1++
		} else if arr1[p2] < arr2[p1] {
			p2++
		}
	}
	return res
}
