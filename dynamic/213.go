package dynamic

import "fmt"

type Test213 struct {
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans1 := rob(nums[1:])

	ans2 := rob(nums[:len(nums)-1])

	return max(ans1, ans2)
}

func (T Test213) Run() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob2(nums))
}
