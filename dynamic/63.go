package dynamic

import "fmt"

type Test63 struct{}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	x := len(obstacleGrid)
	y := len(obstacleGrid[0])
	var dp [][]int = make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, y)
	}

	for i := 0; i < y; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
			for j := i + 1; j < y; j++ {
				dp[0][j] = 0
			}
			break
		} else {
			dp[0][i] = 1
		}
	}
	for i := 0; i < x; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			for j := i + 1; j < x; j++ {
				dp[j][0] = 0
			}
			break
		} else {
			dp[i][0] = 1
		}
	}
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[x-1][y-1]
}

func (T Test63) Run() {
	//var obs [][]int = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	var obs [][]int = [][]int{{1, 0}}
	fmt.Println(uniquePathsWithObstacles(obs))
}
