package dynamic

import "fmt"

type Test322 struct {
}

func coinChange(coins []int, amount int) int {
	var dp []int = make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func (T Test322) Run() {
	coins := []int{2}
	fmt.Println(coinChange(coins, 3))
}
