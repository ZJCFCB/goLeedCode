package binaryTree

import "fmt"

type Test513 struct{}

type answer struct {
	val   int
	hight int
}

func findleft(root *TreeNode, ans *answer, hight int) {
	if root == nil {
		return
	}
	findleft(root.Left, ans, hight+1)
	findleft(root.Right, ans, hight+1)

	if hight > ans.hight {
		ans.hight = hight
		ans.val = root.Val
	}

}

func findBottomLeftValue(root *TreeNode) int {
	var ans *answer = &answer{val: 0, hight: 0}
	findleft(root, ans, 1)
	return ans.val
}

/*
//广度优先搜索
func findBottomLeftValue(root *TreeNode) (ans int) {
    q := []*TreeNode{root}
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node.Right != nil {
            q = append(q, node.Right)
        }
        if node.Left != nil {
            q = append(q, node.Left)
        }
        ans = node.Val
    }
    return
}
*/

func (T Test513) Run() {
	nums := []int{1, 2, 3, 4, -1, 5, 6, -1, -1, -1, -1, 7}
	t := NewTree(nums)
	fmt.Println(PrintTree(t, 2))
	fmt.Println(findBottomLeftValue(t))
}
