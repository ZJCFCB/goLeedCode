package hot100

import "fmt"

type Test51 struct{}

func solveNQueens(n int) [][]string {
	var ans [][]string
	var dpy, dpdj, dpxdj []bool = make([]bool, n), make([]bool, n*2), make([]bool, n*2)
	var temp [][]byte
	for i := 0; i < n; i++ {
		var str []byte
		for j := 0; j < n; j++ {
			str = append(str, '.')
		}
		temp = append(temp, str)
	}

	var dfs func(x int)

	dfs = func(x int) {
		if x == n {
			var anstemp []string
			for i := 0; i < len(temp); i++ {
				anstemp = append(anstemp, string(temp[i]))
			}
			ans = append(ans, anstemp)
			return
		}
		for y := 0; y < n; y++ {
			if !dpy[y] && !dpdj[y-x+n] && !dpxdj[x+y] {
				dpy[y] = true
				dpdj[y-x+n] = true
				dpxdj[x+y] = true
				temp[x][y] = 'Q'
				dfs(x + 1)
				temp[x][y] = '.'
				dpy[y] = false
				dpdj[y-x+n] = false
				dpxdj[x+y] = false
			}
		}
	}

	dfs(0)
	return ans
}

func (T Test51) Run() {
	fmt.Println(solveNQueens(4))
}
