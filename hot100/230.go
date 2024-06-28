package hot100

import "fmt"

type Test230 struct{}

func kthSmallest(root *TreeNode, k int) int {
	var ans []int
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		if len(ans) == k {
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans[k-1]
}
func (T Test230) Run() {
	root := NewTree([]int{3, 1, 4, -1, 2})
	x := kthSmallest(root, 2)
	fmt.Println(x)
}
