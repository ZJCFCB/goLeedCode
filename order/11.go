package order

import (
	"fmt"
)

type Test11 struct{}

func maxArea(height []int) int {
	var ans int
	var left, right int
	left = 0
	right = len(height) - 1
	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func (T Test11) Run() {

	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(nums))
}
