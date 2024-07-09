package classic

import "fmt"

type Test80 struct{}

func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func (T Test80) Run() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates1(nums)
	fmt.Println(nums)
}
