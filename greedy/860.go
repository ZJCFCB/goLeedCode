package greedy

import "fmt"

type Test860 struct{}

func lemonadeChange(bills []int) bool {
	var count [3]int
	for _, v := range bills {
		fmt.Println(count)
		switch v {
		case 5:
			count[0]++
		case 10:
			if count[0] == 0 {
				return false
			} else {
				count[1]++
				count[0]--
			}
		case 20:
			if count[1] > 0 {
				if count[0] > 0 {
					count[1]--
					count[0]--
					count[2]++
				} else {
					return false
				}
			} else {
				if count[0] >= 3 {
					count[0] -= 3
					count[2]++
				} else {
					return false
				}
			}

		}
	}
	return true
}

func (T Test860) Run() {
	var ratings []int = []int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}
	fmt.Println(lemonadeChange(ratings))
}
