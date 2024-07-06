package hot100

import "fmt"

type Test73 struct{}

func setZeroes(matrix [][]int) {
	lenx, leny := len(matrix), len(matrix[0])
	var flag1, flag2 bool
	for i := 0; i < leny; i++ {
		if matrix[0][i] == 0 {
			flag1 = true
		}
	}

	for i := 0; i < lenx; i++ {
		if matrix[i][0] == 0 {
			flag2 = true
		}
	}

	for i := 1; i < lenx; i++ {
		for j := 1; j < leny; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < lenx; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < leny; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < leny; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < lenx; j++ {
				matrix[j][i] = 0
			}
		}
	}
	if flag1 {

		for i := 0; i < leny; i++ {
			matrix[0][i] = 0
		}
	}

	if flag2 {
		for i := 0; i < lenx; i++ {
			matrix[i][0] = 0
		}
	}
}

func (T Test73) Run() {
	m := [][]int{{1}, {0}}
	setZeroes(m)
	fmt.Println(m)
}
