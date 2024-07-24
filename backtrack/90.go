package backtrack

import (
	"fmt"
	"sort"
)

type Test90 struct{}

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var temp []int
	ans = append(ans, []int{})
	sort.Ints(nums)
	var dfs func(start int)
	dfs = func(start int) {
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

func (T Test90) Run() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 3}))
}
