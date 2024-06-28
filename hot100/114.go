package hot100

type Test114 struct{}

// func flatten(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	var dfs func(*TreeNode) (*TreeNode, *TreeNode)

// 	dfs = func(root *TreeNode) (*TreeNode, *TreeNode) {
// 		if root.Left == nil && root.Right == nil {
// 			return root, root
// 		}

// 		var lh, lt, rh, rt *TreeNode
// 		if root.Left != nil {
// 			lh, lt = dfs(root.Left)
// 		}
// 		if root.Right != nil {
// 			rh, rt = dfs(root.Right)
// 		}
// 		if root.Left != nil && root.Right != nil {

// 			root.Left = nil
// 			root.Right = lh
// 			lh.Left = nil
// 			lt.Left = nil
// 			lt.Right = rh
// 			rh.Left = nil
// 			return root, rt
// 		}
// 		if root.Left != nil && root.Right == nil {
// 			root.Left = nil
// 			root.Right = lh
// 			lh.Left = nil
// 			return root, lt
// 		}
// 		if root.Left == nil && root.Right != nil {
// 			root.Left = nil
// 			rh.Left = nil
// 			return root, rt
// 		}
// 		return root, root
// 	}
// 	dfs(root)
// }

func flatten(root *TreeNode) {
	var findright func(*TreeNode) *TreeNode
	findright = func(root *TreeNode) *TreeNode {
		if root.Right == nil {
			return root
		}
		return findright(root.Right)
	}

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			templeft := root.Left
			tempright := root.Right
			allright := findright(templeft)
			root.Left = nil
			root.Right = templeft
			allright.Right = tempright
		}
		dfs(root.Right)
	}
	dfs(root)
}

func (T Test114) Run() {
	root := NewTree([]int{1})
	flatten(root)
	//fmt.Println(root.Right.Right.Val)
}
