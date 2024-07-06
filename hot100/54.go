package hot100

import "fmt"

type Test54 struct{}

func spiralOrder(matrix [][]int) []int {
	var ans []int
	var flag [][]bool = make([][]bool, len(matrix)+2)
	for i := 0; i < len(matrix)+2; i++ {
		flag[i] = make([]bool, len(matrix[0])+2)
	}
	var lenx, leny int = len(matrix), len(matrix[0])
	for i := 0; i < leny+2; i++ {
		flag[0][i] = true
		flag[lenx+1][i] = true
	}

	for i := 0; i < lenx+2; i++ {
		flag[i][0] = true
		flag[i][leny+1] = true
	}
	var x, y int = 0, 0
	var length int
	var work int = 0
	var post [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for {
		if flag[x+1][y+1] == false {
			flag[x+1][y+1] = true
			ans = append(ans, matrix[x][y])
			length++
		}
		if length >= lenx*leny {
			break
		}
		if flag[x+post[work][0]+1][y+post[work][1]+1] == true { //撞墙了
			for i := 0; i < 3; i++ {
				work = (work + 1) % 4
				if flag[x+post[work][0]+1][y+post[work][1]+1] == false {
					break
				}
			}
		}
		x += post[work][0]
		y += post[work][1]
	}
	return ans
}
func (T Test54) Run() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(m))
}
