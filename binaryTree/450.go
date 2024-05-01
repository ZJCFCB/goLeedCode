package binaryTree

import "fmt"

type Test450 struct{}

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	//找到了
	default:
		switch {
		case root.Left == nil && root.Right == nil: //这里，只是返回的nil  并没有实际销毁该节点
			return nil
		case root.Left != nil && root.Right == nil:
			return root.Left
		case root.Right != nil && root.Left == nil:
			return root.Right
		default: //左右都不为空
			//右子树的最左端作为新的根节点
			newRoot := root.Right
			for newRoot.Left != nil {
				newRoot = newRoot.Left
			}
			newRoot.Right = deleteNode(root.Right, newRoot.Val)
			newRoot.Left = root.Left
			return newRoot
		}
	}
	return root
}

func (T Test450) Run() {
	nums1 := []int{5, 3, 6, 2, 4, -1, 7}
	t1 := NewTree(nums1)
	t := deleteNode(t1, 3)
	fmt.Println(PrintTree(t, 1))
}
