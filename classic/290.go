package classic

import (
	"fmt"
	"strings"
)

type Test290 struct{}

func wordPattern(pattern string, s string) bool {
	news := strings.Split(s, " ")
	if len(pattern) != len(news) {
		return false
	}
	var count map[byte]string = make(map[byte]string)
	var countstring map[string]bool = make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		value, ok := count[pattern[i]]
		if ok {
			if value != news[i] {
				return false
			}
		} else {
			if countstring[news[i]] {
				return false
			}
			count[pattern[i]] = news[i]
			countstring[news[i]] = true
		}
	}
	return true
}

func (T Test290) Run() {
	fmt.Println(wordPattern("abbc", "a b b c"))
}
