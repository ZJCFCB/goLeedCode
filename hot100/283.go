package hot100

import "fmt"

type Test283 struct{}

func moveZeroes(nums []int) {
	var first, second int = 0, 0
	length := len(nums)

	for first <= second && first < length && second < length {
		for first < length && nums[first] != 0 {
			first++
		}
		second = first + 1
		for second < length && nums[second] == 0 {
			second++
		}
		if second < length {
			nums[second], nums[first] = nums[first], nums[second]
		}
	}
	fmt.Println(nums)
}

func (T Test283) Run() {
	strs := []int{0, 1, 0, 3, 12}
	moveZeroes(strs)
}
