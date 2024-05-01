package binaryTree

import "fmt"

type Test669 struct{}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	if root.Val < low || root.Val > high { //执行移除操作
		return deleteNode(root, root.Val) //直接移除吧，，，懒得写了
	}

	return root
}

/*
好好看看，，，，什么是优雅，，不复杂。注意审题。。
不要把问题想复杂了，，，根本就没必要挨个删掉阿。。
因为题目没有要求，，，修剪后的树仍然是一颗AVL树。。。

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}




*/

func (T Test669) Run() {
	nums1 := []int{3, 0, 4, -1, 2, -1, -1, -1, -1, 1}
	t1 := NewTree(nums1)
	t := trimBST(t1, 1, 3)
	fmt.Println(PrintTree(t, 1))
}
