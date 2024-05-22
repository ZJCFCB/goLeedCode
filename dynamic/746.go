package dynamic

import "fmt"

type Test746 struct{}

/*
从小到大和从大到小遍历都是一样的
*/
func minCostClimbingStairs(cost []int) int {
	var dp []int = make([]int, len(cost)+2)
	dp[len(cost)-1] = cost[len(cost)-1]
	dp[len(cost)-2] = cost[len(cost)-2]
	for i := len(cost) - 3; i > 0; i-- {
		dp[i] = min(dp[i+1], dp[i+2]) + cost[i]
	}
	fmt.Println(dp)
	return min(dp[1], dp[2]+cost[0])
}

func (T Test746) Run() {
	var cost []int = []int{10, 15}
	fmt.Println(minCostClimbingStairs(cost))
}
