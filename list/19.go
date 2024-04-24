package list

import (
	"fmt"
)

type Test19 struct{}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ahead, behind, newHead *ListNode
	newHead = &ListNode{Next: head}
	ahead = head

	for i := 0; i < n; i++ {
		ahead = ahead.Next
	}

	behind = newHead

	for ahead != nil {
		ahead = ahead.Next
		behind = behind.Next
	}

	behind.Next = behind.Next.Next

	return newHead.Next
}

func (test Test19) Run() {
	array := []int{1}
	head := NewListByArray(array)
	result := removeNthFromEnd(head, 1)
	ans := TurnListToArray(result)
	fmt.Println(ans)
}
