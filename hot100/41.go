package hot100

import "fmt"

type Test41 struct{}

func abs(nums int) int {
	if nums < 0 {
		return -nums
	}
	return nums
}
func firstMissingPositive(nums []int) int {
	var length int = len(nums)
	//先排除不可能的,且保证全部>0
	for i := 0; i < length; i++ {
		if nums[i] >= length+1 || nums[i] <= 0 {
			nums[i] = length + 1
		}
	}

	for i := 0; i < length; i++ {
		post := abs(nums[i])
		if post < length+1 {
			nums[post-1] = -abs(nums[post-1])
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}
func (T Test41) Run() {
	str1 := []int{1, 2, 3, 4}
	x := firstMissingPositive(str1)
	fmt.Println(x)
}
