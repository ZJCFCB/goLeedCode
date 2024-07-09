package classic

import "fmt"

type Test88 struct{}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var flag1, flag2, work int = m - 1, n - 1, m + n - 1
	for flag1 >= 0 && flag2 >= 0 {
		if nums1[flag1] > nums2[flag2] {
			nums1[work] = nums1[flag1]
			flag1--
		} else {
			nums1[work] = nums2[flag2]
			flag2--
		}
		work--
	}
	for work >= 0 {
		if flag1 < 0 {
			nums1[work] = nums2[flag2]
			flag2--
		} else {
			nums1[work] = nums1[flag1]
			flag1--
		}
		work--
	}
}
func (T Test88) Run() {
	nums1 := make([]int, 5)
	nums1[0] = 1
	nums1[1] = 2
	nums1[2] = 3

	nums2 := []int{4, 5}
	merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
}
