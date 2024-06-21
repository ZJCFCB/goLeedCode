package order

import (
	"fmt"
)

type Test438 struct{}

func findAnagrams(s string, p string) []int {

	var result []int
	slen, plen := len(s), len(p)
	if slen < plen {
		return result
	}

	var sCount, pCount [26]int
	for i := 0; i < plen; i++ {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}
	if sCount == pCount {
		result = append(result, 0)
	}
	for i := plen; i < slen; i++ {
		sCount[s[i-plen]-'a']--
		sCount[s[i]-'a']++
		if sCount == pCount {
			result = append(result, i-plen+1)
		}
	}
	return result
}

func (T Test438) Run() {
	str1 := "cbaebabacd"
	str2 := "abc"
	fmt.Println(findAnagrams(str1, str2))
}
