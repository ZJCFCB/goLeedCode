package dynamic

import "fmt"

type Test337 struct{}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	selected := root.Val + left[1] + right[1]
	notselected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, notselected}
}

func rob3(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func (T Test337) Run() {
	root := []int{3, 2, 3, -1, 3, -1, 1}
	t := NewTree(root)
	fmt.Println(rob3(t))
}
