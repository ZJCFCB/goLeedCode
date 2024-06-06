package monotonicstack

import "fmt"

type Test42 struct{}

func trap(height []int) int {
	var result int

	var stack []int = make([]int, 0)
	for k, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] { // 大于栈顶元素，需要出栈计算
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			result += (k - left - 1) * (min(height[left], v) - height[top])
		}
		stack = append(stack, k)
	}
	return result
}

func (T Test42) Run() {
	nusm2 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(nusm2))
}
