package binaryTree

import "fmt"

type Test700 struct{}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	} else {
		if val > root.Val {
			return searchBST(root.Right, val)
		} else {
			return searchBST(root.Left, val)
		}
	}
}

func (T Test700) Run() {
	nums1 := []int{4, 2, 7, 1, 3}
	t1 := NewTree(nums1)
	t := searchBST(t1, 9)
	fmt.Println(PrintTree(t, 1))
}
