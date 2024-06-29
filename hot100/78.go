package hot100

import "fmt"

type Test78 struct{}

func subsets(nums []int) [][]int {
	var ans [][]int
	var temp []int
	var dfs func(int)

	ans = append(ans, []int{})

	dfs = func(post int) {
		if post >= len(nums) {
			//ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := post; i < len(nums); i++ {
			temp = append(temp, nums[i])
			ans = append(ans, append([]int{}, temp...))
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

func (T Test78) Run() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
