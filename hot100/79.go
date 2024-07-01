package hot100

import "fmt"

type Test79 struct{}

func exist(board [][]string, word string) bool {
	lenx, leny := len(board), len(board[0])
	var dfs func(int, int, int) bool
	var path [][]bool = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		path[i] = make([]bool, len(board[0]))
	}
	dfs = func(i, j int, post int) bool {
		fmt.Println(path)
		if post == len(word) {
			return true
		}
		p := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		for b := 0; b < 4; b++ {
			x := i + p[b][0]
			y := j + p[b][1]
			if x >= 0 && x < lenx && y >= 0 && y < leny && !path[x][y] && board[x][y] == string(word[post]) {
				path[x][y] = true
				if dfs(x, y, post+1) {
					return true
				}
				path[x][y] = false
			}
		}

		return false
	}
	for i := 0; i < lenx; i++ {
		for j := 0; j < leny; j++ {
			if board[i][j] == string(word[0]) {
				fmt.Println(i, j)
				path[i][j] = true
				if dfs(i, j, 1) {
					return true
				}
				path[i][j] = false
			}
		}
	}
	return false
}

func (T Test79) Run() {
	board := [][]string{
		{"A", "A"},
	}
	fmt.Println(exist(board, "AAA"))
}
