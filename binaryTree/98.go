package binaryTree

import "fmt"

type Test98 struct{}

func post(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = post(root.Left, result)
	result = append(result, root.Val)
	result = post(root.Right, result)
	return result
}
func isValidBST(root *TreeNode) bool {
	var result []int
	var ans bool = true
	result = post(root, result)
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			ans = false
			break
		}
	}
	return ans
}

func (T Test98) Run() {
	nums1 := []int{2, 1, 3}
	t1 := NewTree(nums1)
	t := isValidBST(t1)
	fmt.Println(t)
}
