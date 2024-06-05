package dynamic

import "fmt"

type Test198 struct{}

func rob(nums []int) int {
	var dp []int = make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func (T Test198) Run() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
