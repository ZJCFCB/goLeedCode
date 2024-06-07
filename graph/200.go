package graph

import "fmt"

type Test200 struct{}

func numIslands(grid [][]string) int {
	var lenx int = len(grid)
	var leny int = len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = "0"
		if x-1 >= 0 && grid[x-1][y] == "1" {
			dfs(x-1, y)
		}
		if x+1 < lenx && grid[x+1][y] == "1" {
			dfs(x+1, y)
		}
		if y-1 >= 0 && grid[x][y-1] == "1" {
			dfs(x, y-1)
		}
		if y+1 < leny && grid[x][y+1] == "1" {
			dfs(x, y+1)
		}
	}
	var num int
	for x := 0; x < lenx; x++ {
		for y := 0; y < leny; y++ {
			if grid[x][y] == "1" {
				num++
				dfs(x, y)
			}
		}
	}
	return num
}

func (T Test200) Run() {
	var graph [][]string = [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	}
	fmt.Println(numIslands(graph))

}
