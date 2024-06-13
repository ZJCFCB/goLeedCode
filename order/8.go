package order

import (
	"fmt"
	"strconv"
	"strings"
)

type Test8 struct{}

func zhenfu(s string) (bool, bool) {
	if s[0] == '-' {
		return false, true
	}
	if s[0] == '+' {
		return true, true
	}
	return true, false
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	zf, ok := zhenfu(s)
	if ok == true {
		s = s[1:]
	}
	var result int
	for i := 0; i < len(s); i++ {
		n := s[i]
		x, err := strconv.Atoi(string(n))
		if err != nil {
			break
		}
		result = result*10 + x
		if zf && result >= 1<<31-1 {
			return 1<<31 - 1
		}
		if !zf && -result < -1<<31 {
			return -1 << 31
		}
	}
	if zf {
		return result
	}
	return -result
}
func (T Test8) Run() {

	fmt.Println(myAtoi("    words and 987    "))
}
