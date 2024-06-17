package order

import (
	"fmt"
)

type Test14 struct{}

func longestCommonPrefix(strs []string) string {
	var result int
	var flag bool = true
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for flag {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == 0 {
				return ""
			}

			if result >= len(strs[i]) || result >= len(strs[i-1]) || strs[i-1][result] != strs[i][result] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		result++
	}
	return strs[0][:result]
}

func (T Test14) Run() {
	strs := []string{"a", "ac"}
	fmt.Println(longestCommonPrefix(strs))
}
