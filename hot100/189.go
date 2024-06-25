package hot100

import "fmt"

type Test189 struct{}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
func (T Test189) Run() {
	str1 := []int{1, 2}
	rotate(str1, 3)
	fmt.Println(str1)
}
