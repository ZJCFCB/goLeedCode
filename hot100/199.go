package hot100

import "fmt"

type Test199 struct{}

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var queuq []*TreeNode
	queuq = append(queuq, root)
	for len(queuq) != 0 {
		length := len(queuq)
		ans = append(ans, queuq[length-1].Val)
		for i := 0; i < length; i++ {
			temp := queuq[0]
			queuq = queuq[1:]
			if temp.Left != nil {
				queuq = append(queuq, temp.Left)
			}
			if temp.Right != nil {
				queuq = append(queuq, temp.Right)
			}
		}
	}
	return ans
}

func (T Test199) Run() {
	root := NewTree([]int{1, 2, 3, -1, 5, -1, 4})
	x := rightSideView(root)
	fmt.Println(x)
}
