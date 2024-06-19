package order

import (
	"fmt"
)

type Test26 struct{}

func removeDuplicates(nums []int) int {
	var count map[int]bool = make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := count[nums[i]]
		if ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			count[nums[i]] = true
		}
	}
	fmt.Println(nums)
	return len(nums)
}

// func removeDuplicates(nums []int) int {
// 	uniqCursor := 0

// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[uniqCursor] {
// 			continue
// 		}

// 		uniqCursor++
// 		if i != uniqCursor {
// 			nums[uniqCursor] = nums[i]
// 		}
// 	}

// 	return uniqCursor + 1
// }

func (T Test26) Run() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
