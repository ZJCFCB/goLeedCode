package string

import "fmt"

type Test459 struct{}

func ifEqual(s string, step int) bool {
	var length int = len(s)
	for begin := 0; begin < step; begin++ {
		for j := begin + step; j < length; j += step {
			if s[j] != s[j-step] {
				return false
			}
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return true
	}
	var step []int = make([]int, 0)
	var length = len(s)
	for i := 1; i < length; i++ {
		if s[i] == s[0] && length%i == 0 { //把不是长度的整数倍的排除掉
			step = append(step, i)
		}
	}
	for _, v := range step {
		if ifEqual(s, v) == true {
			return true
		}
	}

	return false
}

/*

解答牛逼。。。。

func repeatedSubstringPattern(s string) bool {
	str := s + s
	return strings.Contains(str[1:len(str)-1], s)
}

*/

func (T Test459) Run() {
	var s string = "a"
	fmt.Println(repeatedSubstringPattern(s))
}
