package binaryTree

import "fmt"

type Test538 struct{}

func countTree(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	countTree(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	countTree(root.Left, sum)
	return root.Val
}
func convertBST(root *TreeNode) *TreeNode {
	var sum *int = new(int)
	countTree(root, sum)
	return root
}

func (T Test538) Run() {
	nums1 := []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}
	t1 := NewTree(nums1)
	t := convertBST(t1)
	fmt.Println(PrintTree(t, 1))
}

/*
func countTree(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	sum = countTree(root.Right, sum)
	sum += root.Val
	root.Val = sum

	return countTree(root.Left, sum)
}
func convertBST(root *TreeNode) *TreeNode {
	countTree(root, 0)
	return root
}
*/
