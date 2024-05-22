package dynamic

import "fmt"

type Test509 struct{}

func fib(n int) int {
	var dp []int = make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func (T Test509) Run() {
	fmt.Println(fib(0))
}
