package greedy

import "fmt"

type Test45 struct{}

func jump(nums []int) int {
	var dp []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = len(nums) + 1
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= nums[i]+i; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}

func (T Test45) Run() {
	nums := []int{2}
	fmt.Println(jump(nums))
}
