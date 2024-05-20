package greedy

import "fmt"

type Test134 struct {
}

func canCompleteCircuit(gas []int, cost []int) int {
	var rem []int = make([]int, len(gas))

	for i := 0; i < len(gas); i++ {
		rem[i] = gas[i] - cost[i]
	}
	if len(rem) == 1 {
		if rem[0] >= 0 {
			return 0
		}
		return -1
	}
	var count int
	for i := 0; i < len(gas); i++ {
		if rem[i] > 0 {
			count = 0
			for j := i; j < len(gas); j++ {
				count += rem[j]
				if count < 0 {
					break
				}
			}
			if count >= 0 {
				for j := 0; j < i; j++ {
					count += rem[j]
					if count < 0 {
						break
					}
				}
				if count >= 0 {
					return i
				}
			}
		}
	}
	return -1
}

func (T Test134) Run() {
	var gas []int = []int{2}
	var cost []int = []int{3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
