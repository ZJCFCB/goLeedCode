package hot100

import "fmt"

type Test139 struct{}

func wordBreak(s string, wordDict []string) bool {
	var wordmap map[string]bool = make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordmap[v] = true
	}
	var dp []bool = make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordmap[s[j:i]] {
				dp[i] = true
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]
}

func (T Test139) Run() {
	dic := []string{"leed", "code"}
	s := "leedcode"
	fmt.Println((wordBreak(s, dic)))
}
