package backtrack

import (
	"fmt"
	"sort"
)

type Test47 struct{}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	var st []bool = make([]bool, len(nums))
	var dfs func(post int)
	sort.Ints(nums)
	dfs = func(cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && !st[i-1] { // 去重，用st来判别是深度还是广度
				continue
			}
			if !st[i] {
				path = append(path, nums[i])
				st[i] = true
				dfs(cur + 1)
				st[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func (T Test47) Run() {
	nums := []int{1, 2, 3}
	fmt.Println(permuteUnique(nums))
}
