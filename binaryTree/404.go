package binaryTree

import "fmt"

type Test404 struct{}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	} else {
		//判断一下左边是否为叶子
		if root.Left.Left == nil && root.Left.Right == nil {
			return root.Left.Val + sumOfLeftLeaves(root.Right)
		}
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

/*

使用后序遍历要简介一些
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)
	rightValue := sumOfLeftLeaves(root.Right)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}
	return leftValue + rightValue
}


*/

func (T Test404) Run() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	t := NewTree(nums)
	fmt.Println(sumOfLeftLeaves(t))
}
