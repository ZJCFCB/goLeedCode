package graph

import "fmt"

type Test685 struct{}

func findRedundantDirectedConnection(edges [][]int) []int {
	var father []int = make([]int, 1001)
	for i := 0; i < 1001; i++ {
		father[i] = i
	}
	var find func(v int) int
	var ans []int = make([]int, 2)
	find = func(v int) int {
		if father[v] == v {
			return v
		} else {
			father[v] = find(father[v])
			return father[v]
		}
	}

	var issame func(u, v int) bool = func(u, v int) bool {
		u = find(u)
		v = find(v)
		return u == v
	}

	var join func(u, v int) = func(u, v int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		} else {
			father[u] = father[v]
		}
	}

	var indegree []int = make([]int, len(edges)+1) // n个节点，n条边,这是题目的意思
	for i := 0; i < len(edges); i++ {
		indegree[edges[i][1]]++
	}
	var vec []int //记录入度为2的节点
	for i := len(edges) - 1; i >= 0; i-- {
		if indegree[edges[i][1]] == 2 {
			vec = append(vec, i)
		}
	}

	var isTreeAfterDelete func(delete int) bool = func(delete int) bool {
		for i := 0; i < 1001; i++ {
			father[i] = i
		}
		for i := 0; i < len(edges); i++ {
			if i == delete {
				continue
			}
			if issame(edges[i][0], edges[i][1]) {
				return false
			}
			join(edges[i][0], edges[i][1])
		}
		return true
	}

	var removeEdges func() []int = func() []int {
		for i := 0; i < 1001; i++ {
			father[i] = i
		}
		for i := 0; i < len(edges); i++ {
			if issame(edges[i][0], edges[i][1]) {
				return []int{edges[i][0], edges[i][1]}
			}
			join(edges[i][0], edges[i][1])
		}
		return []int{}
	}
	if len(vec) > 0 { //处理有入度>2的节点的情况
		if isTreeAfterDelete(vec[0]) {
			ans[0] = edges[vec[0]][0]
			ans[1] = edges[vec[0]][1]
		} else {
			ans[0] = edges[vec[1]][0]
			ans[1] = edges[vec[1]][1]
		}
	} else {
		ans = removeEdges() //处理没有入度>2的节点的情况，此时代表存在环
	}

	return ans
}

func (T Test685) Run() {
	grid := [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}
	fmt.Println(findRedundantDirectedConnection(grid))
}
