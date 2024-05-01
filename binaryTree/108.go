package binaryTree

import "fmt"

type Test108 struct{}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	root := &TreeNode{
		Val:   nums[middle],
		Left:  nil,
		Right: nil,
	}
	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle+1:])
	return root
}

func (T Test108) Run() {
	num := []int{-10, -3, 0, 5, 9}

	x := sortedArrayToBST(num)
	fmt.Println(PrintTree(x, 1))
}
