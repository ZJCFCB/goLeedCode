package hot100

import (
	"fmt"
)

type Test102 struct{}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int = make([][]int, 0)
	if root == nil {
		return ans
	}
	var queue []*TreeNode = make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		tempans := []int{}
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			tempans = append(tempans, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		ans = append(ans, tempans)
	}
	return ans
}

func (T Test102) Run() {
	root := NewTree([]int{1, 2, 3, 4, 5})
	fmt.Println(levelOrder(root))
}
