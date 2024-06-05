package dynamic

import "fmt"

type Test377 struct{}

func combinationSum4(nums []int, target int) int {
	var dp []int = make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] > i {
				continue
			} else {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
func (T Test377) Run() {
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum4(nums, 4))
}
