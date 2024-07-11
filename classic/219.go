package classic

import "fmt"

type Test219 struct{}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var count map[int]int = make(map[int]int)
	for index, v := range nums {
		value, ok := count[v]
		if ok && abs(index, value) <= k { //存在
			fmt.Println(value, index)
			return true
		}
		count[v] = index
	}
	return false
}

func (T Test219) Run() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
