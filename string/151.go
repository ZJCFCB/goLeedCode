package string

import "fmt"

type Test151 struct{}

func reverseWords(s string) string {
	var tempResult []string = make([]string, 0)
	var flag int = len(s) - 1
	var result string

	for flag >= 0 {
		if s[flag] == ' ' {
			flag--
			continue
		} else { // 不等于空格
			var work int = flag
			for work >= 0 && s[work] != ' ' {
				work--
			}
			tempResult = append(tempResult, string(s[work+1:flag+1]))
			flag = work
		}
	}

	for i := 0; i < len(tempResult); i++ {
		if i == len(tempResult)-1 {
			result += tempResult[i]
		} else {
			result += tempResult[i] + " "
		}
	}
	return result
}

func (T Test151) Run() {
	var s string = "             hello          world         dfd   "
	fmt.Println(reverseWords(s))
}
