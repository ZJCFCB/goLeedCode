package hot100

import "fmt"

type Test994 struct{}

func orangesRotting(grid [][]int) int {
	var flag bool = true
	var ans int
	var lenx, leny int = len(grid), len(grid[0])
	var post [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for flag {
		var tempgrid [][]int = make([][]int, len(grid))
		for i := 0; i < len(grid); i++ {
			tempgrid[i] = make([]int, len(grid[i]))
			copy(tempgrid[i], grid[i])
		}
		flag = false
		for i := 0; i < len(tempgrid); i++ {
			for j := 0; j < len(tempgrid[i]); j++ {
				if tempgrid[i][j] == 2 {
					for x := 0; x < 4; x++ {
						tempx, tempy := i+post[x][0], j+post[x][1]
						if tempx >= 0 && tempx < lenx && tempy >= 0 && tempy < leny && tempgrid[tempx][tempy] == 1 {
							grid[tempx][tempy] = 2
							flag = true
						}
					}
					grid[i][j] = 0
				}
			}
		}
		if flag {
			ans++
		}
	}
	for i := 0; i < lenx; i++ {
		for j := 0; j < leny; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}

func (T Test994) Run() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
