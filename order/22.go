package order

import (
	"fmt"
)

type Test22 struct{}

func generateParenthesis(n int) []string {
	var result []string
	var temp string
	var f func(open, close int)

	f = func(open int, close int) {
		if len(temp) == n*2 {
			a := temp
			result = append(result, a)
			return
		}
		if open < n {
			temp += "("
			f(open+1, close)
			temp = temp[:len(temp)-1]
		}
		if close < open {
			temp += ")"
			f(open, close+1)
			temp = temp[:len(temp)-1]
		}
	}
	f(0, 0)
	return result
}

func (T Test22) Run() {
	x := generateParenthesis(3)
	fmt.Println(x[len(x)-1])
}
