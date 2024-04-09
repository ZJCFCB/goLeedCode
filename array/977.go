package array

import (
	"fmt"
)

type Test977 struct {
}

func sortedSquares(nums []int) []int {
	var left, right, end int = 0, 0, len(nums) - 1
	var answer []int
	for right < end && nums[right] <= 0 {
		right++ //找到正负数的分界点
	}
	left = right - 1

	for left >= 0 && right <= end {
		if -nums[left] > nums[right] {
			answer = append(answer, nums[right]*nums[right])
			right++
		} else {
			answer = append(answer, nums[left]*nums[left])
			left--
		}
	}

	if left < 0 { //负数计算完了，还剩下正数
		for right <= end {
			answer = append(answer, nums[right]*nums[right])
			right++
		}
	}

	if right > end {
		for left >= 0 {
			answer = append(answer, nums[left]*nums[left])
			left--
		}
	}
	return answer
}

func (test Test977) Run() {
	nums := []int{0}
	ans := sortedSquares(nums)

	fmt.Println(ans)
}
