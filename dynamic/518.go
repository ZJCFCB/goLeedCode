package dynamic

import "fmt"

type Test518 struct{}

func change(amount int, coins []int) int {
	var dp []int = make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < len(dp); j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func (T Test518) Run() {
	coins := []int{1, 2, 5}
	fmt.Println(change(5, coins))
}
