package hot100

import (
	"fmt"
)

type Test105 struct{}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	r := preorder[0]
	root := &TreeNode{Val: r}
	if len(preorder) == 1 {
		return root
	}
	var preLeft, preRight, inLeft, inRight []int
	var flag int
	for inorder[flag] != r {
		flag++
	}
	inLeft = inorder[:flag]
	inRight = inorder[flag+1:]
	preLeft = preorder[1 : 1+len(inLeft)]
	preRight = preorder[1+len(inLeft):]
	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, inRight)
	return root
}
func (T Test105) Run() {
	pre := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(pre, inorder)
	fmt.Println(PrintTree(root, 3))
}
