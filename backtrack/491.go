package backtrack

import (
	"fmt"
)

type Test491 struct{}

var (
	res  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	dfs(nums, 0)
	return res
}
func dfs(nums []int, start int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
	for i := start; i < len(nums); i++ {
		if used[nums[i]] { // 去重
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
}

func (T Test491) Run() {
	fmt.Println(findSubsequences([]int{1, 2, 6, 7, 7}))
}
