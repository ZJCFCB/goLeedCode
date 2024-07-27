package graph

import (
	"fmt"
)

type Test695 struct{}

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	var post [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var bfs func(int, int) int = func(x, y int) int {
		var total int
		var queue [][2]int = make([][2]int, 0)
		queue = append(queue, [2]int{x, y})
		total++
		grid[x][y] = 0
		for len(queue) > 0 {
			tx, ty := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := 0; i <= 3; i++ {
				ttx, tty := tx+post[i][0], ty+post[i][1]
				if ttx >= 0 && ttx < len(grid) && tty >= 0 && tty < len(grid[0]) && grid[ttx][tty] == 1 {
					queue = append(queue, [2]int{ttx, tty})
					total++
					grid[ttx][tty] = 0

				}
			}
		}
		return total
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				x := bfs(i, j)
				ans = max(x, ans)
			}
		}
	}
	return ans
}

func (T Test695) Run() {
	var grid [][]int = [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}
