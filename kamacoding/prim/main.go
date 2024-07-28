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
			graph[i][j] = 10001
		}
	}

	var from, to, value int
	for i := 0; i < e; i++ {
		fmt.Scanln(&from, &to, &value)
		graph[from][to] = value
		graph[to][from] = value
	}

	var minDist []int = make([]int, v+1) //记录最小距离
	for i := 0; i <= v; i++ {
		minDist[i] = 10001
	}
	var visited []bool = make([]bool, v+1) //记录是否在生成树里
	var allvisit func() bool = func() bool {
		for i := 1; i <= v; i++ {
			if visited[i] == false {
				return false
			}
		}
		return true
	}
	for !allvisit() {
		//第一步，选一个距离最近的点
		var cur, min int = -1, int(^uint(0) >> 1)
		for j := 1; j <= v; j++ {
			if !visited[j] && minDist[j] < min {
				min = minDist[j]
				cur = j
			}
		}
		//加入生成树
		visited[cur] = true
		//更新距离
		for j := 1; j <= v; j++ {
			if !visited[j] && minDist[j] > graph[cur][j] {
				minDist[j] = graph[cur][j]
			}
		}
	}
	var result int
	for i := 2; i <= v; i++ {
		result += minDist[i]
	}
	fmt.Println(result)
}
