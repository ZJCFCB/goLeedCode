package stackAndqueue

import (
	"fmt"
	"strconv"
)

type Test150 struct{}

func getResult(num1, num2 int, sign string) int {
	switch sign {
	case "+":
		return num1 + num2
	case "-":
		return num2 - num1
	case "*":
		return num1 * num2
	case "/":
		return num2 / num1
	}
	return 0
}

func evalRPN(tokens []string) int {
	var stack []int = make([]int, 0, len(tokens))
	for _, v := range tokens {
		num, ok := strconv.Atoi(v)
		if ok == nil {
			stack = append(stack, num) //如果是数字，直接入栈
		} else {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := getResult(num1, num2, v)
			stack = append(stack, result)
		}
	}
	return stack[len(stack)-1]
}

func (T Test150) Run() {
	var s []string = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(s))
}
