package binaryTree

import "fmt"

type Test226 struct {
}

func (T Test226) Run() {
	num := []int{1, 2, 3, -1, 4, 5, 6, -1, -1, -1, -1, -1, -1, 7, -1}
	t := NewTree(num)

	fmt.Println(PrintTree(t, 3))
}
