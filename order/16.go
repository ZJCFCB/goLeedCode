package order

import (
	"fmt"
	"math"
	"sort"
)

type Test16 struct{}

func threeSumClosest(nums []int, target int) int {
	var result int = math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func (T Test16) Run() {
	nums := []int{0, 3, 97, 102, 200}
	fmt.Println(threeSumClosest(nums, 300))
}
