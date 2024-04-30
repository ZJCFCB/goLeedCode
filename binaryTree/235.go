package binaryTree

import "fmt"

type Test235 struct{}

func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root {
		return p
	}
	if q == root {
		return q
	}

	left := lowestCommonAncestor_(root.Left, p, q)
	right := lowestCommonAncestor_(root.Right, p, q)

	if (left == p && right == q) || (left == q && right == p) {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func (T Test235) Run() {
	num := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	t := NewTree(num)
	t = lowestCommonAncestor_(t, t.Left, t.Right)

	fmt.Println(t.Val)
}
