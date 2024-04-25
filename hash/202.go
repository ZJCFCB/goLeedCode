package hash

import "fmt"

type Test202 struct{}

func isHappy(n int) bool {
	var count map[int]bool = make(map[int]bool, 100)
	var work int = n
	for {
		if work == 1 {
			return true
		} else {
			_, ok := count[work]
			if ok {
				return false
			} else {
				count[work] = true
				work = Getsum(work)
			}
		}
	}
}

func Getsum(n int) int {
	var total int = 0
	for n != 0 {
		f := n % 10
		total += f * f
		n = n / 10
	}
	return total
}

func (T Test202) Run() {
	fmt.Println(isHappy(1999999999))
}
