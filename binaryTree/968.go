package binaryTree

type Test968 struct{}

func minCameraCover(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == 0 || right == 0 { //有一个没覆盖，必须放一个摄像头
			ans++
			return 1
		}
		if left == 1 || right == 1 { //孩子存在摄像头，直接返回有覆盖
			return 2
		}
		return 0
	}
	x := dfs(root)
	//还要检查一下根节点.
	if x == 0 {
		ans++
	}
	return ans
}

func (T Test968) Run() {}
