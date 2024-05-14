package greedy

import (
	"fmt"
	"sort"
)

type Test455 struct{}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var result int = 0
	var flag int = len(s) - 1
	for flag >= 0 {
		if len(g) == 0 {
			break
		} else {
			if s[flag] >= g[len(g)-1] { // 可以发
				result++
				flag--
			}
			g = g[:len(g)-1]
		}
	}
	return result
}

func (T Test455) Run() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}
