package hot100

import (
	"fmt"
)

type Test543 struct{}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		temp := left + right
		ans = max(ans, temp)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}

func (T Test543) Run() {
	root := NewTree([]int{1, 2, 3, 4, 5})
	fmt.Println(diameterOfBinaryTree(root))
}
