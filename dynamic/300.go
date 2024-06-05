package dynamic

import "fmt"

type Test300 struct{}

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var dp []int = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	var result int

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if result <= dp[i] {
			result = dp[i]
		}
	}

	return result
}

func (T Test300) Run() {
	nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(nums))
}
