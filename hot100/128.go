package hot100

import (
	"fmt"
)

type Test128 struct{}

// func longestConsecutive(nums []int) int {
// 	sort.Ints(nums)
// 	var ans, temp int = 1, 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			continue
// 		}
// 		if nums[i]-nums[i-1] == 1 {
// 			temp++
// 			ans = max(ans, temp)
// 		} else {
// 			temp = 1
// 		}
// 	}
// 	return ans
// }

func longestConsecutive(nums []int) int {
	var count map[int]bool = make(map[int]bool)
	for _, v := range nums {
		count[v] = true
	}
	var ans int
	for i := 0; i < len(nums); i++ {
		if count[nums[i]] {
			count[nums[i]] = false
			var temp = 1
			for z := nums[i] - 1; ; z-- {
				_, ok := count[z]
				if ok {
					temp++
					count[z] = false
				} else {
					break
				}
			}
			for z := nums[i] + 1; ; z++ {
				_, ok := count[z]
				if ok {
					temp++
					count[z] = false
				} else {
					break
				}
			}
			ans = max(ans, temp)
		}
	}
	return ans
}

func (T Test128) Run() {
	strs := []int{0, 1, 12, 12, 12, 13, 13}
	fmt.Println(longestConsecutive(strs))
}
