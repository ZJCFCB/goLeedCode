package order

import "fmt"

type Test7 struct{}

func reverse(x int) int {
	var ans int
	var flag bool = true
	if x < 0 {
		flag = false
		x = -x
	}
	for x != 0 {
		ans = ans*10 + x%10
		if ans > 1<<31-1 {
			return 0
		}
		x = x / 10
	}
	if flag {
		return ans
	} else {
		return -ans
	}
}

func (T Test7) Run() {

	fmt.Println(reverse(120))
}
