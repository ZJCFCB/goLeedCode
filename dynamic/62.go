package dynamic

import "fmt"

type Test62 struct{}

func uniquePaths(m int, n int) int {
	var dp [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//初始化第一行和第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func (T Test62) Run() {
	fmt.Println(uniquePaths(3, 2))
}
