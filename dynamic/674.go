package dynamic

import "fmt"

type Test674 struct{}

func findLengthOfLCIS(nums []int) int {
	var dp []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	var result int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > result {
				result = dp[i]
			}
		}
	}
	return result
}
func (T Test674) Run() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}
