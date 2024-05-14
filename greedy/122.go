package greedy

import "fmt"

type Test122 struct{}

func maxProfit(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}
func (T Test122) Run() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}
