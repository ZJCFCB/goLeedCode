package dynamic

import "fmt"

type Test279 struct{}

func numSquares(n int) int {
	var dp []int = make([]int, n+1)
	var sq []int = make([]int, 0)

	for i := 1; i <= n; i++ {
		sq = append(sq, i*i)
	}
	for i := 0; i < n+1; i++ {
		dp[i] = n + 2
	}

	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < len(sq); j++ {
			if i >= sq[j] {
				dp[i] = min(dp[i], dp[i-sq[j]]+1)
			}
		}
	}
	return dp[n]
}

func (T Test279) Run() {
	fmt.Println(numSquares(13))
}
