package binaryTree

import "fmt"

type Test101 struct{}

func leftRight(left, right *TreeNode) bool {
	if left == nil {
		if right == nil {
			return true
		} else {
			return false
		}
	} else {
		if right == nil {
			return false
		} else {
			if left.Val == right.Val {
				return leftRight(left.Left, right.Right) && leftRight(left.Right, right.Left)
			} else {
				return false
			}
		}
	}
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return leftRight(root.Left, root.Right)
	}
}

func (T Test101) Run() {
	nums := []int{1, 2, 2, -1, 3, -1, 3}
	t := NewTree(nums)
	fmt.Println(isSymmetric(t))
}
