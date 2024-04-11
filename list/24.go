package list

import (
	"fmt"
)

type Test24 struct{}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead *ListNode = &ListNode{Val: 0, Next: head}
	var behind, middle, forhead *ListNode = newHead, newHead.Next, newHead.Next.Next
	for {
		temp := forhead.Next
		behind.Next = forhead
		forhead.Next = middle
		middle.Next = temp
		if temp == nil || temp.Next == nil {
			break
		} else {
			behind = behind.Next.Next
			middle = behind.Next
			forhead = middle.Next
		}
	}
	return newHead.Next
}

func (test Test24) Run() {
	array := []int{}
	head := NewListByArray(array)
	result := swapPairs(head)
	ans := TurnListToArray(result)
	fmt.Println(ans)
}
