package hot100

import "fmt"

type Test32 struct{}

func longestValidParentheses(s string) int {
	var stack []int
	var dp []bool = make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		if s[i] == ')' {
			if s[stack[len(stack)-1]] == '(' {
				dp[i] = true
				dp[stack[len(stack)-1]] = true
				stack = stack[:len(stack)-1]
			}
		}
	}
	var tempans, ans int
	for i := 0; i < len(dp); i++ {
		tempans = 0
		for i < len(dp) && dp[i] {
			tempans++
			i++
		}
		ans = max(ans, tempans)
	}
	return ans
}

func (T Test32) Run() {
	str := "()(((())"
	fmt.Println(longestValidParentheses(str))
}
