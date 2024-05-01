package string

import "fmt"

type Test344 struct{}

func reverseString(s []byte) {
	var first, second int = 0, len(s) - 1

	for first < second {
		s[first], s[second] = s[second], s[first]
		first++
		second--
	}
	fmt.Println(string(s))
}

func (T Test344) Run() {
	var s []byte = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
}
