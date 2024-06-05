package dynamic

import "fmt"

type Test1143 struct{}

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp [][]int = make([][]int, len(text2))
	for i := 0; i < len(text2); i++ {
		dp[i] = make([]int, len(text1))
	}
	var result int

	for i := 0; i < len(text1); i++ {
		if text1[i] == text2[0] {
			result = 1
			for j := i; j < len(text1); j++ {
				dp[0][j] = 1
			}
			break
		}
	}

	for i := 0; i < len(text2); i++ {
		if text1[0] == text2[i] {
			result = 1
			for j := i; j < len(text2); j++ {
				dp[j][0] = 1
			}
			break
		}
	}

	for i := 1; i < len(text2); i++ {
		for j := 1; j < len(text1); j++ {
			if text1[j] == text2[i] {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j-1]+1)
				if dp[i][j] > result {
					result = dp[i][j]
				}
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return result
}
func (T Test1143) Run() {
	str1 := "ezupkr"
	str2 := "ubmrapg"
	fmt.Println(longestCommonSubsequence(str1, str2))
}
