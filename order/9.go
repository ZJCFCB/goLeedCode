package order

import (
	"fmt"
	"strconv"
)

type Test9 struct{}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	var first, end int = 0, len(s) - 1
	for first < end {
		if s[first] != s[end] {
			return false
		}
		first++
		end--
	}
	return true
}

func (T Test9) Run() {
	s := -121
	fmt.Println(isPalindrome(s))
}
