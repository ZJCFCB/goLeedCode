package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// 根据一个层次遍历的切片，生成一颗二叉树，并返回二叉树的根节点
// 默认二叉树的数值均为正数
// 在切片data中，使用-1 代表nil
func NewTree(data []int) *TreeNode {
	var treeSet []*TreeNode = make([]*TreeNode, 0, len(data)+1)

	treeSet = append(treeSet, NewNode(len(data))) // 添加一个头节点，放在第0号位置

	for _, v := range data {
		if v == -1 {
			treeSet = append(treeSet, nil)
		} else {
			treeSet = append(treeSet, NewNode(v))
		}
	}

	var length int = len(data) + 1
	for i := 1; i < length; i++ {
		if treeSet[i] == nil {
			continue
		}
		if i*2 < length {
			treeSet[i].Left = treeSet[i*2]
		} else {
			treeSet[i].Left = nil
		}

		if i*2+1 < length {
			treeSet[i].Right = treeSet[i*2+1]
		} else {
			treeSet[i].Right = nil
		}
	}

	return treeSet[1]
}

func PrintTree(T *TreeNode, choice int) []int { // 1 表示先序遍历 2 标识中序遍历  3 表示后续遍历
	var result []int = make([]int, 0, 10)
	switch choice {
	case 1:
		result = Preorder(T, result)
	case 2:
		result = Inorder(T, result)
	case 3:
		result = Postorder(T, result)
	}
	return result
}
func Preorder(T *TreeNode, result []int) []int {
	result = append(result, T.Val)
	if T.Left != nil {
		result = Preorder(T.Left, result)
	}
	if T.Right != nil {
		result = Preorder(T.Right, result)
	}
	return result
}

func Inorder(T *TreeNode, result []int) []int {
	if T.Left != nil {
		result = Inorder(T.Left, result)
	}
	result = append(result, T.Val)
	if T.Right != nil {
		result = Inorder(T.Right, result)
	}
	return result
}

func Postorder(T *TreeNode, result []int) []int {
	if T.Left != nil {
		result = Postorder(T.Left, result)
	}
	if T.Right != nil {
		result = Postorder(T.Right, result)
	}
	result = append(result, T.Val)
	return result
}
