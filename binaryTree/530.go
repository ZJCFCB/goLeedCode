package binaryTree

import "fmt"

type Test530 struct{}

func post_(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = post(root.Left, result)
	result = append(result, root.Val)
	result = post(root.Right, result)
	return result
}

func getMinimumDifference(root *TreeNode) int {
	var result []int
	result = post_(root, result)
	var ans int = int(^uint(0) >> 1) //最大的int
	for i := 1; i < len(result); i++ {
		ans = min(ans, result[i]-result[i-1])
	}
	return ans
}

func (T Test530) Run() {
	nums1 := []int{1, 0, 48, -1, -1, 12, 49}
	t1 := NewTree(nums1)
	t := getMinimumDifference(t1)
	fmt.Println(t)
}
