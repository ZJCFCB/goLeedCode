package hot100

import (
	"fmt"
)

type Test437 struct{}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var ans int
	var dfs func(*TreeNode) map[int]int
	dfs = func(root *TreeNode) map[int]int {
		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				ans++
			}
			return map[int]int{root.Val: 1}
		}
		var leftmap, rightmap map[int]int = make(map[int]int), make(map[int]int)
		if root.Left != nil {
			leftmap = dfs(root.Left)
		}
		if root.Right != nil {
			rightmap = dfs(root.Right)
		}

		var mymap map[int]int = make(map[int]int)
		mymap[root.Val]++
		if root.Val == targetSum {
			ans++
		}
		for k, v := range leftmap {
			if root.Val+k == targetSum {
				ans += v
			}
			mymap[root.Val+k] += v
		}
		for k, v := range rightmap {
			if root.Val+k == targetSum {
				ans += v
			}
			mymap[root.Val+k] += v
		}
		return mymap
	}
	dfs(root)
	return ans
}

/*
前缀和

	func pathSum(root *TreeNode, targetSum int) (ans int) {
	    preSum := map[int64]int{0: 1}
	    var dfs func(*TreeNode, int64)
	    dfs = func(node *TreeNode, curr int64) {
	        if node == nil {
	            return
	        }
	        curr += int64(node.Val)
	        ans += preSum[curr-int64(targetSum)]
	        preSum[curr]++
	        dfs(node.Left, curr)
	        dfs(node.Right, curr)
	        preSum[curr]--
	        return
	    }
	    dfs(root, 0)
	    return
	}
*/
func (T Test437) Run() {
	root := NewTree([]int{1, -2, -3, 1, 3, -2, -100, -1})
	fmt.Println(pathSum(root, -1))
}
