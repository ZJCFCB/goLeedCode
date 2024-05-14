package greedy

import "fmt"

type Test376 struct{}

//一定一定要注意，处理相等的情况

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	var dp []int = make([]int, len(nums))
	var flag int
	dp[0] = 1
	if nums[1] == nums[0] {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	if nums[1]-nums[0] > 0 {
		flag = 1
	} else if nums[1]-nums[0] < 0 {
		flag = 2
	} else {
		flag = 3
	}

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 {
			if flag == 2 || flag == 3 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1]
			}
			flag = 1
		}
		if nums[i]-nums[i-1] < 0 {
			if flag == 1 || flag == 3 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1]
			}
			flag = 2
		}
		if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-1]
}

func (T Test376) Run() {
	nums := []int{3, 2}
	fmt.Println(wiggleMaxLength(nums))
}
