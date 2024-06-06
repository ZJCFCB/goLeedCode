package monotonicstack

import "fmt"

type Test739 struct{}

func dailyTemperatures(temperatures []int) []int {
	var stack []int = make([]int, 0)
	var ans []int = make([]int, len(temperatures))
	stack = append(stack, 0)

	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 {
			j := stack[len(stack)-1]               //栈顶元素
			if temperatures[i] > temperatures[j] { // 大于栈顶元素
				ans[j] = i - j
				stack = stack[:len(stack)-1] // 出栈
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	return ans
}

func (T Test739) Run() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
}
