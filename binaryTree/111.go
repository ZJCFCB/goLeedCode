package binaryTree

import "fmt"

type Test111 struct{}

func mdep(T *TreeNode, nowDeep int, ans *int) {
	if T.Left == nil && T.Right == nil {
		if *ans > nowDeep {
			*ans = nowDeep
		}
		return
	}
	if T.Left != nil {
		mdep(T.Left, nowDeep+1, ans)
	}
	if T.Right != nil {
		mdep(T.Right, nowDeep+1, ans)
	}
}

func minDepth(root *TreeNode) int {
	var ans *int = new(int) //指针，一定要给分配空间
	if root == nil {
		return 0
	}
	*ans = 1000
	mdep(root, 1, ans)
	return *ans
}

/*

看看别人的优雅的代码

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
    if root.Left==nil && root.Right==nil {
        return 1
    }
    if root.Left==nil {
        return minDepth(root.Right)+1
    }
    if root.Right==nil {
        return minDepth(root.Left)+1
    }
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

*/

func (T Test111) Run() {
	num := []int{3, 9, 20, -1, -1, 15, 7}
	t := NewTree(num)
	x := minDepth(t)

	fmt.Println(x)
}
