package dynamic

import "fmt"

type Test139 struct{}

func wordBreak(s string, wordDict []string) bool {
	var wordDictSet map[string]bool = make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordDictSet[v] = true
	}

	var dp []bool = make([]bool, len(s)+1)

	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func (T Test139) Run() {
	s := "leedcode"
	wordDict := []string{"leed", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
