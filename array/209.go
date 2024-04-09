package array

import (
	"fmt"
)

type Test209 struct{}

func minSubArrayLen(target int, nums []int) int {
	var left, right, sum, ans int = 0, 0, 0, len(nums) + 1

	for right < len(nums) || sum >= target { //一定要注意边界值的处理
		if sum < target {
			sum += nums[right]
			right++
		} else {
			for sum >= target {
				sum -= nums[left]
				left++
			}
			if right-left+1 < ans {
				ans = right - left + 1
			}
		}
	}

	if ans == len(nums)+1 {
		ans = 0
	}
	return ans
}

func (test Test209) Run() {
	nums := []int{2, 3, 1, 2, 4, 3}
	ans := minSubArrayLen(7, nums)

	fmt.Println(ans)
}
