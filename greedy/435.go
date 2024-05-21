package greedy

import (
	"fmt"
	"sort"
)

type Test435 struct{}

func eraseOverlapIntervals(intervals [][]int) int {
	var ans int
	//这里一定要注意一下排序的方式
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	//fmt.Println(intervals)
	var maxRight = intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		//fmt.Println(maxRight)
		if intervals[i][0] < maxRight {
			ans++
		} else {
			maxRight = intervals[i][1]
		}
	}
	return ans
}

func (T Test435) Run() {
	var people [][]int = [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}
	fmt.Println(eraseOverlapIntervals(people))
}
