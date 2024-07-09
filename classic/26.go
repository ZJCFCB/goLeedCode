package classic

import "fmt"

type Test26 struct{}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var flag int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[flag] {
			continue
		}
		flag++
		nums[flag] = nums[i]
	}
	return flag + 1
}
func (T Test26) Run() {
	nums := []int{1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6, 7}
	x := removeDuplicates(nums)
	fmt.Println(x)
}
