package array

import (
	"fmt"
)

type Test27 struct{}

func removeElement(nums []int, val int) int {
	var left, right int
	left = 0

	right = len(nums) - 1

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return right + 1 //right一直指向数组的最尾端
}

func (test Test27) Run() {
	nums := []int{0}
	ans := removeElement(nums, 0)

	fmt.Println(ans)
}
