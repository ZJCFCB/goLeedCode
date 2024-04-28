package binaryTree

import "fmt"

type Test226 struct {
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

func (T Test226) Run() {
	num := []int{4, 2, 7, 1, 3, 6, 9}
	t := NewTree(num)
	t = invertTree(t)

	fmt.Println(PrintTree(t, 1))
}
