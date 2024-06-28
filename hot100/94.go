package hot100

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var mt func(*TreeNode)

	mt = func(root *TreeNode) {
		if root == nil {
			return
		}
		mt(root.Left)
		ans = append(ans, root.Val)
		mt(root.Right)
	}
	mt(root)
	return ans
}
