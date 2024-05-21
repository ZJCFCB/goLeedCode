package greedy

import (
	"fmt"
	"strconv"
)

type Test738 struct{}

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	//找到出现递减的地方
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] { // 往回走，这个循环主要针对 = 的情况
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}

func (T Test738) Run() {
	fmt.Println(monotoneIncreasingDigits(12345386))
}
