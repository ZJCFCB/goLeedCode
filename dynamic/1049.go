package dynamic

import "fmt"

type Test1049 struct{}

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	var max, sum, target, length int
	length = len(stones)
	var dp [][]bool = make([][]bool, length)
	for i := 0; i < length; i++ {
		sum += stones[i]
		if stones[i] > max {
			max = stones[i]
		}
	}
	target = sum/2 + 1
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, target+1)
	}
	var ans int

	for i := 0; i < length; i++ {
		dp[i][0] = true
	}
	dp[0][stones[0]] = true
	ans = stones[0]

	for i := 1; i < length; i++ {
		v := stones[i]
		for j := 1; j < target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j-v] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			if dp[i][j] && j > ans {
				ans = j
			}
		}
	}
	return sum - 2*ans
}

func (T Test1049) Run() {
	nums := []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(nums))

}
