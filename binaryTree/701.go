package binaryTree

import "fmt"

type Test701 struct{}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val < val { //往右走
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func (T Test701) Run() {
	nums1 := []int{4, 2, 7, 1, 3}
	t1 := NewTree(nums1)
	t := insertIntoBST(t1, 5)
	fmt.Println(PrintTree(t, 1))
}
