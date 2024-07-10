package classic

import "fmt"

func lengthOfLastWord(s string) int {
	var flag, ans int = len(s) - 1, 0
	for flag >= 0 && s[flag] == ' ' {
		flag--
	}
	for flag >= 0 && s[flag] != ' ' {
		flag--
		ans++
	}
	return ans
}

type Test58 struct{}

func (T Test58) Run() {
	fmt.Println(lengthOfLastWord("a"))
}
