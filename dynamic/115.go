package dynamic

import "fmt"

type Test115 struct{}

func numDistinct(s string, t string) int {
	var dp [][]int = make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		dp[0][i] = 1
	}
	fmt.Println(dp)
	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
		if i == 1 {
			fmt.Println(dp)
		}
	}
	return dp[len(t)][len(s)]
}

func (T Test115) Run() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
