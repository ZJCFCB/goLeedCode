package list

import (
	"fmt"
)

type Test203 struct{}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	var forhead, behind, newHead *ListNode
	newHead = &ListNode{0, head}
	forhead = newHead.Next
	behind = newHead
	for forhead != nil {
		if forhead.Val == val {
			behind.Next = forhead.Next
			forhead = behind.Next
		} else {
			forhead = forhead.Next
			behind = behind.Next
		}
	}
	return newHead.Next
}

func (test Test203) Run() {
	array := []int{1, 2, 6, 3, 4, 5, 6}
	head := NewListByArray(array)
	result := removeElements(head, 7)
	ans := TurnListToArray(result)
	fmt.Println(ans)

}
