package order

import (
	"fmt"
)

type Test28 struct{}

func strStr(haystack string, needle string) int {
	var next []int = make([]int, len(needle))
	next[0] = 0
	var j int
	for i := 1; i < len(needle); i++ {
		//先进行j的回溯
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[j] == needle[i] {
			j++
		}
		next[i] = j
	}
	j = 0
	for i := 0; i < len(haystack); i++ {
		//先决定匹配从哪里开始
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func (T Test28) Run() {
	str1 := "sasasasasadbutsad"
	str2 := "sad"
	fmt.Println(strStr(str1, str2))
}
