package backtrack

import "fmt"

type Test216 struct{}

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var temp []int
	var f func(int, int)

	f = func(count int, flag int) {
		fmt.Println(temp, count, flag)
		if flag > 10 || count > n || len(temp) > k {
			return
		}
		if count == n && len(temp) == k {
			var tempre []int = make([]int, k)
			copy(tempre, temp)

			result = append(result, tempre)
			return
		}

		temp = append(temp, flag)
		count = count + flag
		f(count, flag+1)
		temp = temp[:len(temp)-1]
		count = count - flag
		f(count, flag+1)

	}
	f(0, 1)
	return result
}

/*
看看别人的优雅的代码

var res [][]int
var path []int

	func combinationSum3(k int, n int) [][]int {
		res, path = make([][]int, 0), make([]int, 0, k)
		dfs(n, k, 1, 0)
		return res
	}

	func dfs(n, k, start, sum int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i <= 9; i++ {
			if sum+i > n || 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(n, k, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
*/
func (T Test216) Run() {
	result := combinationSum3(9, 45)
	fmt.Println(result)
}
