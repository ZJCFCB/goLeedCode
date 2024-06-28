package hot100

import "fmt"

type Test46 struct{}

func permute(nums []int) [][]int {
	var ans [][]int
	var temp []int
	var used []bool = make([]bool, len(nums))

	var dfs func(post int)
	length := len(nums)
	dfs = func(post int) {
		if len(temp) == length {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := 0; i < length; i++ {
			if used[i] {
				continue
			}
			temp = append(temp, nums[i])
			used[i] = true
			dfs(i + 1)
			temp = temp[:len(temp)-1]
			used[i] = false
		}
	}
	dfs(0)
	return ans
}

func (T Test46) Run() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
