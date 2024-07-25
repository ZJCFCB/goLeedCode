package dynamic

import (
	"fmt"
)

type Test471 struct{}

func findMaxForm(strs []string, m, n int) int {
	var length int = len(strs)
	var num0, num1 []int = make([]int, length), make([]int, length)
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				num0[i]++
			} else {
				num1[i]++
			}
		}
	}
	var dp [][]int = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		//dp[i][0] = 1
	}
	// for i := 0; i <= n; i++ {
	// 	dp[0][i] = 1
	// }
	fmt.Println(dp)
	for x := 0; x < length; x++ {
		for y := m; y >= num0[x]; y-- {
			for z := n; z >= num1[x]; z-- {
				dp[y][z] = max(dp[y][z], dp[y-num0[x]][z-num1[x]]+1)

			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (T Test471) Run() {
	str := []string{"10", "0", "1"}
	fmt.Println(findMaxForm(str, 1, 1))
}
