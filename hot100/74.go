package hot100

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

type Test74 struct{}

func (T Test74) Run() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//matrix := [][]int{{1}}
	fmt.Println(searchMatrix(matrix, 1))
}
