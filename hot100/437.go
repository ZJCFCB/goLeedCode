package hot100

import (
	"fmt"
)

type Test437 struct{}

func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		if root.Val == targetSum {
			ans++
		}
		dfs(root.Left, targetSum)
		dfs(root.Left, targetSum-root.Val)
		dfs(root.Right, targetSum)
		dfs(root.Right, targetSum-root.Val)
	}
	dfs(root, targetSum)
	return ans
}
func (T Test437) Run() {
	root := NewTree([]int{1, 2, 3, 4, 5})
	fmt.Println(pathSum(root, 4))
}
