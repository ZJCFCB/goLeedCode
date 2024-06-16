package order

import (
	"fmt"
)

type Test10 struct{}

func isMatch(s string, p string) bool {
	var first, second int
	for first < len(s) && second < len(p) {
		if p[second] == '.' {
			first++
			second++
			continue
		}

		if p[second] == '*' {
			wn := p[second-1]
			for first < len(s) && second+1 < len(p) && (wn == '.' || s[first] == wn) && (s[first] == p[second+1]) {
				first++
			}
			second++
			continue
		}
		if s[first] != p[second] {
			second++
		} else {
			first++
			second++
		}
	}
	if first == len(s) && second == len(p) {
		return true
	} else {
		return false
	}
}

func (T Test10) Run() {
	s := "aaa"
	ss := "a*a"

	fmt.Println(isMatch(s, ss))
}
