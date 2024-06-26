package hot100

import "fmt"

type Test238 struct{}

func productExceptSelf(nums []int) []int {
	var l, r []int = make([]int, len(nums)), make([]int, len(nums))
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	var ans []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = l[i] * r[i]
	}
	return ans
}

func (T Test238) Run() {
	str1 := []int{1, 2, 3, 4}
	x := productExceptSelf(str1)
	fmt.Println(x)
}
