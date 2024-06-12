package order

import "fmt"

type Test5 struct{}

func longestPalindrome(s string) string {
	var centor, length int
	var ans string
	for centor = 1; centor < len(s); centor++ {
		length = 1
		for centor-length >= 0 && centor+length < len(s) {
			if s[centor+length] == s[centor-length] {
				length++
			} else {
				break
			}
		}
		if len(ans) < (length-1)*2 {
			ans = s[centor-length+1 : centor+length]
		}
	}

	for centor = 0; centor < len(s)-1; centor++ {
		if s[centor] == s[centor+1] {
			length = 1
			for centor-length >= 0 && centor+length+1 < len(s) {
				if s[centor+length+1] == s[centor-length] {
					length++
				} else {
					break
				}
			}
			if len(ans) < length*2 {
				ans = s[centor-length+1 : centor+length+1]
			}
		}
	}
	if len(ans) == 0 {
		return string(s[0])
	}
	return ans
}

func (T Test5) Run() {
	s := "babbbbbbbbbac"
	fmt.Println(longestPalindrome(s))
}
