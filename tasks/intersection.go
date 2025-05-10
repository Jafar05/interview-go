package tasks

func Intersection(one, two []int) []int {
	m := make(map[int]int)
	var res []int

	for _, val := range one {
		m[val]++
	}

	for _, val := range two {
		if count, ok := m[val]; ok && count > 0 {
			m[val]--
			res = append(res, val)
		}
	}
	return res
}
