package binaryTree

import "fmt"

type Test257 struct{}

func nowPath(root *TreeNode, s string, result []string) []string {

	if root.Left == nil && root.Right == nil { // 走到了叶子节点
		s += fmt.Sprintf("->%d", root.Val)
		result = append(result, s[2:])
	}

	s += fmt.Sprintf("->%d", root.Val)
	if root.Left != nil {
		result = nowPath(root.Left, s, result)
	}
	if root.Right != nil {
		result = nowPath(root.Right, s, result)
	}
	return result
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string = make([]string, 0)
	return nowPath(root, "", result)
}

func (T Test257) Run() {
	nums := []int{1, 2, 3, -1, 5}
	t := NewTree(nums)
	fmt.Println(binaryTreePaths(t))
}
