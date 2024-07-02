package hot100

import "fmt"

type Test64 struct{}

func minPathSum(grid [][]int) int {
	lenx, leny := len(grid), len(grid[0])
	var dp [][]int = make([][]int, lenx)
	for i := 0; i < lenx; i++ {
		dp[i] = make([]int, leny)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < leny; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < lenx; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < lenx; i++ {
		for j := 1; j < leny; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[lenx-1][leny-1]
}

func (T Test64) Run() {
	var grid [][]int = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
