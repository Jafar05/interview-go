package two_pointers

func MergeTwoSlice(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 2

	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}
