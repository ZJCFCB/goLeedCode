package binaryTree

import "fmt"

type Test236 struct{}

// 先遍历一遍，得到节点的父节点的对应关系

func getRoot(root *TreeNode, result map[*TreeNode]*TreeNode) map[*TreeNode]*TreeNode {
	if root == nil {
		return result
	}
	if root.Left != nil {
		result[root.Left] = root
	}

	if root.Right != nil {
		result[root.Right] = root
	}
	result = getRoot(root.Left, result)
	result = getRoot(root.Right, result)
	return result
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var result map[*TreeNode]*TreeNode = make(map[*TreeNode]*TreeNode)
	result[root] = nil
	result = getRoot(root, result)
	// for k, v := range result {
	// 	if v == nil {
	// 		fmt.Printf("k = %d,v = nil\n", k.Val)
	// 		continue
	// 	}
	// 	fmt.Printf("k = %d,v = %d\n", k.Val, v.Val)
	// }
	//然后再对result 进行处理
	var fatherPath map[*TreeNode]bool = make(map[*TreeNode]bool)
	var work *TreeNode = p
	fatherPath[p] = true
	for {
		work = result[work]
		if work == nil {
			break
		}
		fatherPath[work] = true
	}

	var ans *TreeNode = q
	for {
		if ans == nil {
			break
		}
		_, ok := fatherPath[ans]
		if ok {
			break
		}
		ans = result[ans]
	}
	return ans
}

/*

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if (l == p && r == q) || (l == q && r == p) {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

*/

func (T Test236) Run() {
	num := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	t := NewTree(num)
	t = lowestCommonAncestor(t, t.Left, t.Left.Right.Right)

	fmt.Println(t.Val)
}
