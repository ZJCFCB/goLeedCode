package dynamic

import "fmt"

type Test714 struct{}

func maxProfit714(prices []int, fee int) int {
	var dp [][]int = make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}
func (T Test714) Run() {
	fmt.Println(maxProfit714([]int{1, 3, 2, 8, 4, 9}, 2))
}
