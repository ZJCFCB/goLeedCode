package hot100

import "fmt"

type Test35 struct{}

func searchInsert(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		} else {
			left = mid + 1
		}
	}
	return (left+right)>>1 + 1
}

func (T Test35) Run() {
	fmt.Println(searchInsert([]int{1, 3, 4, 6}, 7))
}
