package binaryTree

import "fmt"

type Test104 struct {
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func (T Test104) Run() {
	num := []int{3, -1, 20}
	t := NewTree(num)
	x := maxDepth(t)

	fmt.Println(x)
}
