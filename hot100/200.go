package hot100

import "fmt"

func numIslands(grid [][]string) int {
	var ans int
	xlen, ylen := len(grid), len(grid[0])
	var bfs func(x, y int)

	var post [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	bfs = func(x, y int) {
		if x >= xlen || y > ylen {
			return
		}
		grid[x][y] = "0"
		for i := 0; i < 4; i++ {
			tempx, tempy := x+post[i][0], y+post[i][1]
			if tempx >= 0 && tempx < xlen && tempy >= 0 && tempy < ylen && grid[tempx][tempy] == "1" {
				bfs(tempx, tempy)
			}
		}
	}

	for i := 0; i < xlen; i++ {
		for j := 0; j < ylen; j++ {
			if grid[i][j] == "1" {
				bfs(i, j)
				ans++
			}
		}
	}
	return ans
}

type Test200 struct{}

func (T Test200) Run() {
	grid := [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "1"},
	}
	fmt.Println(numIslands(grid))
}
