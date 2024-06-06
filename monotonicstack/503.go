package monotonicstack

import "fmt"

type Test503 struct{}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, _ := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}
func (T Test503) Run() {
	nums1 := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums1))
}
