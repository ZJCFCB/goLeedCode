package binaryTree

import "fmt"

type Test617 struct{}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil && root2 != nil {
		return root2
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
	}
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func (T Test617) Run() {
	nums1 := []int{1, 3, 2, 5}
	nums2 := []int{2, 1, 3, -1, 4, -1, 7}
	t1 := NewTree(nums1)
	t2 := NewTree(nums2)
	anst := mergeTrees(t1, t2)
	fmt.Println(PrintTree(anst, 1))
}
