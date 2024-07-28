package main

import (
	"fmt"
)

func main() {
	var v, e int
	fmt.Scanln(&v, &e)
	var graph [][]int = make([][]int, v+1)
	for i := 0; i <= v; i++ {
		graph[i] = make([]int, v+1)
		for j := 0; j <= v; j++ {
			graph[i][j] = 10e9
		}
	}

	var from, to, value int
	for i := 0; i < e; i++ {
		fmt.Scanln(&from, &to, &value)
		graph[from][to] = value
	}
	var start, end int = 1, v

	var minDist []int = make([]int, v+1) //记录从源点开始的最小距离
	for i := 0; i <= v; i++ {
		minDist[i] = 10e9
	}
	var visited []bool = make([]bool, v+1) //记录是否被访问过

	minDist[start] = 0
	//visited[start] = true
	for i := 1; i <= v; i++ {
		//第一步，选一个距离最近的点
		var cur, min int = 1, 10e9
		for j := 1; j <= v; j++ {
			if !visited[j] && minDist[j] < min {
				min = minDist[j]
				cur = j
			}
		}
		//标记访问过了
		visited[cur] = true
		//更新距离
		for j := 1; j <= v; j++ {
			if !visited[j] && minDist[j] > minDist[cur]+graph[cur][j] {
				minDist[j] = minDist[cur] + graph[cur][j]
			}
		}
	}
	if minDist[end] == 10e9 {
		fmt.Println(-1)
	} else {
		fmt.Println(minDist[end])
	}
}
