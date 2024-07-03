package hot100

import "fmt"

type Test131 struct{}

func partition(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	f := make([][]bool, n)

	//先初始化为true，即当个字符为回文串
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	//f数组标记的是 i:j的子串，是否为回文字符串
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func (T Test131) Run() {
	fmt.Println(partition("abccbaa"))
}
