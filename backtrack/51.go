package backtrack

import "fmt"

type Test51 struct{}

func isok(qipan [][]byte, i, j int) bool {
	for flag := 0; flag < len(qipan); flag++ {
		if qipan[flag][j] == 'Q' {
			return false
		}
	}
	for flag := 0; flag < len(qipan); flag++ {
		if qipan[i][flag] == 'Q' {
			return false
		}
	}
	for flag1, flag2 := i, j; flag1 >= 0 && flag2 >= 0; flag1, flag2 = flag1-1, flag2-1 {
		if qipan[flag1][flag2] == 'Q' {
			return false
		}
	}
	for flag1, flag2 := i, j; flag1 < len(qipan) && flag2 < len(qipan); flag1, flag2 = flag1+1, flag2+1 {
		if qipan[flag1][flag2] == 'Q' {
			return false
		}
	}

	for flag1, flag2 := i, j; flag1 >= 0 && flag2 < len(qipan); flag1, flag2 = flag1-1, flag2+1 {
		if qipan[flag1][flag2] == 'Q' {
			return false
		}
	}

	for flag1, flag2 := i, j; flag1 < len(qipan) && flag2 > 0; flag1, flag2 = flag1+1, flag2-1 {
		if qipan[flag1][flag2] == 'Q' {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	var result [][]string = make([][]string, 0)
	var qipan [][]byte = make([][]byte, n)
	for i := 0; i < n; i++ {
		qipan[i] = make([]byte, n)
		for j := 0; i < n; j++ {
			qipan[i][j] = '.'
		}
	}
	var f func(x, y int, n int)

	f = func(x, y int, n int) {
		if n == 0 { // 收割
			var tempResult []string = make([]string, n)
			for _, v := range qipan {
				tempResult = append(tempResult, string(v))
			}
		}
		for i := n - 1; i > 0; i-- {
			for j := n - 1; j > 0; j++ {
				if isok(qipan, i, j) {
					qipan[i][j] = 'Q'
					f(i-1, j-1, n-1)
				}
			}
		}
	}
	return result
}

func (T Test51) Run() {
	fmt.Println(solveNQueens(4))
}
