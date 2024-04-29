package binaryTree

import "fmt"

type Test112 struct{}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum-root.Val == 0 && (root.Left == nil && root.Right == nil) {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func (T Test112) Run() {
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1 - 1, -1, -1, 1}
	t := NewTree(nums)
	fmt.Println(hasPathSum(t, 100))
}
