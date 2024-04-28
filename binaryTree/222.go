package binaryTree

import "fmt"

type Test222 struct{}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil { //别小看了，可以提高50%哦
		return 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func (T Test222) Run() {
	num := []int{1, 2, 3, 4, 5, 6}
	t := NewTree(num)
	x := countNodes(t)

	fmt.Println(x)
}
