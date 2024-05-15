package greedy

import (
	"fmt"
	"sort"
)

type Test1005 struct{}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var flag int
	for flag = 0; flag < len(nums) && k > 0 && nums[flag] < 0; flag++ {
		nums[flag] = -nums[flag]
		k--
	}
	var work int
	if k > 0 { //这时候肯定全>0
		if flag >= len(nums) || (flag > 0 && nums[flag-1] < nums[flag]) {
			work = flag - 1
		} else {
			work = flag
		}
		if k%2 == 1 {
			nums[work] = -nums[work]
		}
	}
	var result int
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
func (T Test1005) Run() {
	nums := []int{-4, -2, -3}
	fmt.Println(largestSumAfterKNegations(nums, 4))
}
