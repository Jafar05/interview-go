package tasks

import (
	"fmt"
	"strconv"
)

func GetRangesMain() {
	fmt.Println(getRanges([]int{0, 1, 2, 3, 4, 7, 8, 10})) // "0-4,7-8,10"
	fmt.Println(getRanges([]int{4, 7, 10}))                // "4,7,10"
	fmt.Println(getRanges([]int{2, 3, 8, 9}))
}

func getRanges(values []int) string {
	if len(values) == 0 {
		return ""
	}
	result := ""

	start, prev := values[0], values[0]

	for i := 1; i < len(values); i++ {
		if values[i] == prev+1 {
			prev = values[i]
		} else {
			if start == prev {
				result += strconv.Itoa(start) + ","
			} else {
				result += strconv.Itoa(start) + "–" + strconv.Itoa(prev) + ","
			}
			start = values[i]
			prev = values[i]
		}
	}
	if start == prev {
		result += strconv.Itoa(start)
	} else {
		result += strconv.Itoa(start) + "–" + strconv.Itoa(start)
	}
	return result
}
