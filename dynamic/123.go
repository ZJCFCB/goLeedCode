package dynamic

import "fmt"

type Test123 struct{}

func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
		fmt.Println(buy1, sell1, buy2, sell2, prices[i])
	}
	return sell2
}

func (T Test123) Run() {
	nums := []int{9, 3, 5, 4, 6, 5, 7, 6, 8}
	fmt.Println(maxProfit3(nums))
}
