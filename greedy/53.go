package greedy

import "fmt"

type Test53 struct{}

func maxSubArray(nums []int) int {
	var result int = nums[0]
	var dp []int = make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
			if result < nums[i] {
				result = nums[i]
			}
		} else {
			if dp[i-1]+nums[i] > result {
				result = dp[i-1] + nums[i]
			}
			dp[i] = dp[i-1] + nums[i]
		}
	}
	return result
}

func (T Test53) Run() {
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
}
