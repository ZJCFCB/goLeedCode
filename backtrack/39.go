package backtrack

import (
	"fmt"
)

type Test39 struct{}

func combinationSum(candidates []int, target int) [][]int {
	var temp []int
	var ans [][]int

	var f func(int, int)

	f = func(target int, flag int) {
		if flag == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		f(target, flag+1)

		if target-candidates[flag] >= 0 {
			temp = append(temp, candidates[flag])
			f(target-candidates[flag], flag)
			temp = temp[:len(temp)-1]
		}
	}
	f(target, 0)
	return ans
}

func (T Test39) Run() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
}
