package backtrack

import (
	"fmt"
	"sort"
)

type Test40 struct{}

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var dfs func(target int, flag int)
	var temp []int
	sort.Ints(candidates)
	dfs = func(target int, flag int) {
		if target == 0 { // 收割
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			ans = append(ans, tmp)
			return
		}
		for i := flag; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if i != flag && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])

			dfs(target-candidates[i], i+1)
			temp = temp[:len(temp)-1]

		}
	}
	dfs(target, 0)
	return ans
}

func (T Test40) Run() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))
}
