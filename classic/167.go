package classic

import "fmt"

type Test167 struct{}

func twoSum(numbers []int, target int) []int {
	var ans []int

	var ef func(left, right int, target int) int
	ef = func(left, right int, target int) int {
		var mid int
		for left <= right {
			mid = (left + right) >> 1
			if numbers[mid] == target {
				return mid
			}
			if numbers[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	var left, right int = 0, len(numbers) - 1
	// for right >= 0 && numbers[right] > target {
	// 	right--
	// }
	for left < right {
		temp := ef(left+1, right, target-numbers[left])
		if temp != -1 {
			return []int{left + 1, temp + 1}
		}
		left++
	}
	return ans
}

func (T Test167) Run() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
