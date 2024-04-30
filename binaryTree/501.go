package binaryTree

import "fmt"

type Test501 struct{}

type ans struct {
	fre int
	ans []int
}

func findModePost(root *TreeNode, result map[int]int) map[int]int {
	if root == nil {
		return result
	}

	result[root.Val]++

	result = findModePost(root.Left, result)
	result = findModePost(root.Right, result)
	return result
}

func findMode(root *TreeNode) []int {
	var post map[int]int = make(map[int]int)
	post = findModePost(root, post)
	var result ans = ans{fre: -1}
	for k, v := range post {
		if v > result.fre { //频次更大
			result.fre = v
			result.ans = result.ans[:0]
			result.ans = append(result.ans, k)
		} else if v == result.fre {
			result.ans = append(result.ans, k)
		}
	}
	return result.ans
}

func (T Test501) Run() {
	nums := []int{2}
	t := NewTree(nums)
	fmt.Println(findMode(t))
}
