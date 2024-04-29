package binaryTree

import "fmt"

type Test106 struct{}

// 思路很清晰但是慢
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	//后序遍历的最后一个节点就是当前的根节点
	var rootVal = postorder[len(postorder)-1]
	var tempnode *TreeNode = &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	//然后根据根节点，划分左右子树
	var leftinorder, rightinorder, leftpostorder, rightpostorder []int
	//先划分中序遍历
	for k, v := range inorder {
		if v == rootVal {
			rightinorder = inorder[k+1:]
			break
		}
		leftinorder = append(leftinorder, v)
	}

	var lenLeft int = len(leftinorder) //左子树的节点数
	for i := 0; i < lenLeft; i++ {
		leftpostorder = append(leftpostorder, postorder[i])
	}
	rightpostorder = postorder[lenLeft : len(postorder)-1]

	tempnode.Left = buildTree(leftinorder, leftpostorder)
	tempnode.Right = buildTree(rightinorder, rightpostorder)
	return tempnode
}

func (T Test106) Run() {
	inorder := []int{9}
	postorder := []int{9}
	t := buildTree(inorder, postorder)
	fmt.Println(PrintTree(t, 1))
}
