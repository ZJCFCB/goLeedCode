package array

import (
	"fmt"
)

func search(nums []int, target int) int {
	var middle, left, right int
	left = 0
	right = len(nums) - 1
	middle = (left + right) / 2
	for left <= right {
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
		middle = (left + right) / 2
	}
	return -1
}

func Test() {
	nums := []int{1, 0, 3, 5, 9, 12}
	ans := search(nums, 10)
	fmt.Println(ans)
}
