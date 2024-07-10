package classic

import "fmt"

type Test125 struct{}

func isPalindrome(s string) bool {
	var news []byte = make([]byte, len(s))
	var l, r int = 0, -1
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			r++
			news[r] = s[i] + 'a' - 'A'
		} else if ('a' <= s[i] && s[i] <= 'z') || '0' <= s[i] && s[i] <= '9' {
			r++
			news[r] = s[i]
		}
	}
	for l <= r {
		if news[l] == news[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func (T Test125) Run() {
	fmt.Println(isPalindrome("Abc,ba"))
}
