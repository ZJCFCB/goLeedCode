package hot100

import "fmt"

type Test560 struct{}

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func (T Test560) Run() {
	strs := []int{1, 1, 2, 3}
	ans := subarraySum(strs, 3)
	fmt.Println(ans)
}
