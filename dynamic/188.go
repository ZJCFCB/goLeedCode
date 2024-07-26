package dynamic

import "fmt"

type Test188 struct{}

func maxProfit188(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))
	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	//方法可以参考2次交易的
	//0 初始化，j+1 第一次持有 j+2 第一次不持有
	//0 初始化的时候，奇数次，即持有，应该初始化为-prices[0]
	// j+1 表示持有，1.延续，2.上一次不持有+买入
	// j+2 表示不持有 1.延续 2.上一次持有+卖出
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func (Test188) Run() {
	fmt.Println(maxProfit188(2, []int{3, 2, 6, 5, 0, 3}))
}
