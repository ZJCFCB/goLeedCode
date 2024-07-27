package graph

import "fmt"

type Test463 struct{}

func islandPerimeter(grid [][]int) int {
	post := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var ans int = 0
	// for i := 0; i < len(grid[0]); i++ {
	// 	if grid[0][i] == 1 {
	// 		ans++
	// 	}
	// 	if grid[len(grid)-1][i] == 1 {
	// 		ans++
	// 	}
	// }

	// for i := 0; i < len(grid); i++ {
	// 	if grid[i][0] == 1 {
	// 		ans++
	// 	}
	// 	if grid[i][len(grid[0])-1] == 1 {
	// 		ans++
	// 	}
	// }
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				for z := 0; z < 4; z++ {
					x, y := i+post[z][0], j+post[z][1]
					if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func (T Test463) Run() {
	var grid [][]int = [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}
