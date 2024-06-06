package dynamic

import "fmt"

type Test72 struct{}

func minDistance1(word1 string, word2 string) int {
	var dp [][]int = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 1; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	dp[0][0] = 0

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]

}

func (T Test72) Run() {
	s := "leetcode"
	t := "etco"
	fmt.Println(minDistance1(s, t))
}
