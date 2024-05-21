package greedy

import (
	"fmt"
	"sort"
)

type Test406 struct{}
type point struct {
	value int
	post  int
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		//fmt.Println(ans)
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}

func (T Test406) Run() {
	var people [][]int = [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 1},
	}
	fmt.Println(reconstructQueue(people))
}
