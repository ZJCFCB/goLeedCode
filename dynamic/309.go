package dynamic

import "fmt"

type Test309 struct{}

func maxProfit309(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	status := make([]int, n*4)
	for i := range dp {
		dp[i] = status[:4]
		status = status[4:]
	}
	dp[0][0] = -prices[0]
	/*
		注意这里的状态转化，一共有四个状态
		1.持有  （1.继续持有，2.当天买入（前一天是否为冷冻期
		2.卖出  （前一天是否为冷冻期
		3.当天卖出 （直接卖出
		4.今天冷冻
		初始化：持有状态应该是-price
	*/
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

func (T Test309) Run() {
	fmt.Println(maxProfit309([]int{1, 2, 3, 0, 2}))
}
