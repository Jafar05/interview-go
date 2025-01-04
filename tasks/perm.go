package tasks

import "fmt"

func perm(arr []rune, l, r int) {
	if l == r {
		fmt.Println(string(arr))
	} else {
		for i := l; i <= r; i++ {
			arr[i], arr[l] = arr[l], arr[i]
			perm(arr, l+1, r)
			arr[i], arr[l] = arr[l], arr[i]
		}
	}
}

func ImportPerm() {
	str := []rune("abc")
	perm(str, 0, len(str)-1)
}
