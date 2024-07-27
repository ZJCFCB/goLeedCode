package graph

import "fmt"

type Test684 struct{}

func findRedundantConnection(edges [][]int) []int {
	var path []int = make([]int, 1001)
	for i := 0; i < 1001; i++ {
		path[i] = i
	}
	var find func(n int) int
	find = func(n int) int {
		if path[n] == n {
			return n
		} else {
			path[n] = find(path[n])
			return path[n]
		}
	}
	var join func(u, v int) bool
	join = func(u, v int) bool {
		u = find(u)
		v = find(v)
		if u == v {
			return false
		}
		path[u] = v
		return true
	}

	for i := 0; i < len(edges); i++ {
		ok := join(edges[i][0], edges[i][1])
		if !ok {
			return []int{edges[i][0], edges[i][1]}
		}
	}
	return []int{1, 2}
}

func (T Test684) Run() {
	fmt.Println(findRedundantConnection([][]int{{0, 1}, {1, 2}, {2, 0}}))
}
