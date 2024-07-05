package hot100

import "fmt"

type Test84 struct{}

func largestRectangleArea(heights []int) int {
	var length int = len(heights)
	stack := []int{}
	var left, right []int = make([]int, length), make([]int, length)
	//先求left
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack1 := []int{}

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack1) > 0 && heights[i] <= heights[stack1[len(stack1)-1]] {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack1) == 0 {
			right[i] = length
		} else {
			right[i] = stack1[len(stack1)-1]
		}
		stack1 = append(stack1, i)
	}
	var ans int
	for i := 0; i < len(heights); i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func (T Test84) Run() {
	nusm2 := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(nusm2))
}
