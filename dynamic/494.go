package dynamic

import "fmt"

type Test494 struct{}

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < target || sum < -target {
		return 0
	}

	var length = sum*2 + 1
	var dp [][]int = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, length)
	}

	dp[0][nums[0]+sum] = 1
	dp[0][-nums[0]+sum] += 1

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 0; j < length; j++ {
			if dp[i-1][j] != 0 && j+v < length && j-v >= 0 {
				dp[i][j+v] += dp[i-1][j]
				dp[i][j-v] += dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target+sum]
}

func (t Test494) Run() {
	nums := []int{100}
	fmt.Println(findTargetSumWays(nums, -200))
}
