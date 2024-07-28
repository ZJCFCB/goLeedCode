package main

import (
	"fmt"
	"math"
)

func main() {
	var v, e int
	fmt.Scanln(&v, &e)
	var graph [][3]int

	var from, to, value int
	for i := 0; i < e; i++ {
		fmt.Scanln(&from, &to, &value)
		graph = append(graph, [3]int{from, to, value})
	}

	var minDist []int = make([]int, v+1) //记录从源点开始的最小距离
	for i := 0; i <= v; i++ {
		minDist[i] = math.MaxInt32
	}

	minDist[1] = 0

	for i := 1; i < v; i++ { //对所有的边  进行v-1次松弛
		for j := 0; j < e; j++ {
			x, y, z := graph[j][0], graph[j][1], graph[j][2]
			if minDist[x] != math.MaxInt32 && minDist[y] > minDist[x]+z {
				minDist[y] = minDist[x] + z
			}
		}
	}
	if minDist[v] == math.MaxInt32 {
		fmt.Println("unconnected")
	} else {
		fmt.Println(minDist[v])
	}
}
