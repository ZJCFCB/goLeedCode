package hot100

import "fmt"

type Test994 struct{}

func orangesRotting(grid [][]int) int {
	var flag bool = true
	var ans int
	for flag {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {

			}
		}
	}
	return ans
}

func (T Test994) Run() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
