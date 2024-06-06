package monotonicstack

import "fmt"

type Test496 struct{}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int = make([]int, 0)
	var m map[int]int = make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1] //出栈
		}
		if len(stack) > 0 {
			m[nums2[i]] = nums2[stack[len(stack)-1]] //当前元素的最近的最大值 就是栈顶
		} else {
			m[nums2[i]] = -1
		}
		stack = append(stack, i)
	}
	var ans []int = make([]int, len(nums1))
	for k, v := range nums1 {
		ans[k] = m[v]
	}
	return ans
}

func (T Test496) Run() {
	nums1 := []int{2, 4}
	nusm2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nusm2))
}
