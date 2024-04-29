package binaryTree

import "fmt"

type Test654 struct{}

func findMax(nums []int) (index int) {
	var ank, anv int = -1, -1
	for k, v := range nums {
		if v > anv {
			ank = k
			anv = v
		}
	}
	return ank
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := findMax(nums)
	root := &TreeNode{
		Val:   nums[index],
		Left:  nil,
		Right: nil,
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func (T Test654) Run() {
	inorder := []int{3, 2, 1, 6, 0, 5}
	t := constructMaximumBinaryTree(inorder)
	fmt.Println(PrintTree(t, 1))
}
