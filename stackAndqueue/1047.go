package stackAndqueue

import "fmt"

type Test1047 struct{}

func removeDuplicates(s string) string {
	var stack []rune = make([]rune, 0)
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		if v == stack[len(stack)-1] { // 相等的话，直接出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}

func (T Test1047) Run() {
	var s string = "abbaca"
	fmt.Println(removeDuplicates(s))
}
