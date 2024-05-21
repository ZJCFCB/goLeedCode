package greedy

import (
	"fmt"
	"sort"
)

type Test56 struct{}

func merge(intervals [][]int) [][]int {
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	var flag int
	var left, right int
	for flag < len(intervals) {
		left = intervals[flag][0]
		right = intervals[flag][1]
		for flag += 1; flag < len(intervals); {
			if intervals[flag][0] <= right {
				right = max(right, intervals[flag][1])
				flag++
			} else {
				break
			}
		}
		result = append(result, []int{left, right})
	}
	return result
}

func (T Test56) Run() {
	var people [][]int = [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	fmt.Println(merge(people))
}
