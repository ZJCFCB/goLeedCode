package graph

import "fmt"

type Test797 struct{}

func allPathsSourceTarget(graph [][]int) [][]int {
	var ans [][]int
	var temp []int
	temp = append(temp, 0)
	var n int = len(graph)
	var f func(int)
	f = func(post int) {
		if post == n-1 {
			ans = append(ans, append([]int(nil), temp...)) // 注意这里，要新建一个temp
			return
		}
		for _, v := range graph[post] {
			temp = append(temp, v)
			f(v)
			temp = temp[:len(temp)-1]
		}
	}

	f(0)
	return ans
}

func (T Test797) Run() {
	var graph [][]int = [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))

}
