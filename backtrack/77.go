package backtrack

import "fmt"

type Test77 struct{}

func combine(n int, k int) [][]int {
	var result [][]int
	var temp []int

	var dfs func(cur int)

	dfs = func(cur int) {
		if len(temp)+n-cur+1 < k {
			return
		}
		if len(temp) == k {
			tempre := make([]int, k)
			copy(tempre, temp)
			result = append(result, tempre)
			return
		} else {
			temp = append(temp, cur)
			dfs(cur + 1)
			temp = temp[:len(temp)-1] //一次回退动作
			dfs(cur + 1)
		}
	}
	dfs(1)
	return result
}

func (T Test77) Run() {
	result := combine(4, 2)
	fmt.Println(result)
}
