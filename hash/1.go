package hash

import "fmt"

type Test1 struct{}

func twoSum(nums []int, target int) []int {
	var count map[int]int = make(map[int]int, len(nums))
	for k, v := range nums {
		index, ok := count[target-v]
		if ok {
			return []int{index, k}
		} else {
			count[v] = k
		}
	}
	return []int{1}
}

func (T Test1) Run() {

	nums := []int{3, 3}
	result := twoSum(nums, 6)
	fmt.Println(result)
}
