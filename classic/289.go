package classic

import "fmt"

type Test289 struct{}

func gameOfLife(board [][]int) {
	var tempboard [][]int = make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		tempboard[i] = make([]int, len(board[0]))
		copy(tempboard[i], board[i])
	}
	var post [][]int = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var count int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count = 0
			for x := 0; x < 8; x++ {
				tempi, tempj := i+post[x][0], j+post[x][1]
				if tempi >= 0 && tempi < len(board) && tempj >= 0 && tempj < len(board[0]) && tempboard[tempi][tempj] == 1 {
					count++
				}
			}
			if tempboard[i][j] == 0 {
				if count == 3 {
					board[i][j] = 1
				}
			} else {
				if count < 2 || count > 3 {
					board[i][j] = 0
				}
			}
		}
	}
}

func (T Test289) Run() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}
