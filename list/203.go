package list

import (
	"fmt"
)

type Test203 struct{}

func removeElements(head *ListNode, val int) *ListNode {
	return head
}

func (test Test203) Run() {
	array := []int{1, 2, 6, 3, 4, 5, 6}
	head := NewListByArray(array)
	result := removeElements(head, 6)
	ans := TurnListToArray(result)
	fmt.Println(ans)

}
