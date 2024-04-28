package binaryTree

import "fmt"

type Test110 struct{}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left) + 1
	if leftDepth <= 0 {
		return -1
	}
	rightDepth := depth(root.Right) + 1
	if rightDepth <= 0 || leftDepth-rightDepth >= 2 || leftDepth-rightDepth <= -2 {
		return -1
	}
	return max(leftDepth, rightDepth)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if depth(root) > 0 {
		return true
	} else {
		return false
	}
}

func (T Test110) Run() {
	num := []int{1, 2, 2, 3, 3, -1, -1, 4, 4}
	t := NewTree(num)
	x := isBalanced(t)

	fmt.Println(x)
}
