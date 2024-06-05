package dynamic

import "fmt"

type Test121 struct{}

func maxProfit(prices []int) int {
	var r []int = make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		r[i] = prices[i] - prices[i-1]
	}
	var flag int
	for flag = 1; flag < len(prices); flag++ {
		if r[flag] > 0 {
			break
		}
	}
	if flag == len(prices) { // 不能获利
		return 0
	}

	var ans int
	var count int

	for ; flag < len(r); flag++ {
		count += r[flag]
		if count < 0 {
			count = 0
		} else {
			if count > ans {
				ans = count
			}
		}

	}
	return ans
}

func (T Test121) Run() {
	nums := []int{7, 1, 5, -90, -6, 100}
	fmt.Println(maxProfit(nums))
}
