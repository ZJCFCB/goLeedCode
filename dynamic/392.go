package dynamic

import "fmt"

type Test392 struct{}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	var dp [][]int = make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(t))
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			for j := i; j < len(t); j++ {
				dp[0][j] = 1
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			for j := i; j < len(s); j++ {
				dp[j][0] = 1
			}
		}
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				for k := j; k < len(t); k++ {
					dp[i][k] = dp[i][j]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)-1][len(t)-1] == len(s)
}
func (T Test392) Run() {
	str1 := "abc"
	str2 := "ahbgdc"
	fmt.Println(isSubsequence(str1, str2))
}
