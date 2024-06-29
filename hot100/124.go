package hot100

import (
	"fmt"
)

type Test124 struct{}

func maxPathSum(root *TreeNode) int {
	var ans int = -100000
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		var left, right int = -100000, -100000

		if root.Left != nil {
			left = dfs(root.Left)
		}
		if root.Right != nil {
			right = dfs(root.Right)
		}
		if root.Val >= 0 {
			ans = max(ans, left+root.Val+right, left+root.Val, right+root.Val, root.Val)
		} else {
			ans = max(ans, left, right, root.Val, left+right+root.Val)
		}
		return max(root.Val, root.Val+max(left, right))
	}
	return max(ans, dfs(root))
}

func (T Test124) Run() {
	root := NewTree([]int{-2, 5, -1, 4, -1, -1, -1, 2, -4})
	fmt.Println(maxPathSum(root))
}
