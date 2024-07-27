package graph

import "fmt"

type Test1971 struct{}

func validPath(n int, edges [][]int, source int, destination int) bool {
	var path []int = make([]int, n)
	for i := 0; i < n; i++ {
		path[i] = i
	}
	var find func(v int) int
	find = func(v int) int {
		if path[v] == v {
			return v
		} else {
			path[v] = find(path[v])
			return path[v]
		}
	}
	var insert func(u, v int) = func(u, v int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		}
		path[u] = v
	}

	for i := 0; i < len(edges); i++ {
		insert(edges[i][0], edges[i][1])
	}
	fmt.Println(path)
	return find(source) == find(destination)
}

func (T Test1971) Run() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
}
