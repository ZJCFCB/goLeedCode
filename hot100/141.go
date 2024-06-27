package hot100

import (
	"fmt"
)

type Test141 struct{}

func hasCycle(head *ListNode) bool {
	var fast, slow *ListNode = head, head
	for fast != nil && slow != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || slow == nil {
		return false
	}
	fast = head
	for fast != head {
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func (test Test141) Run() {
	list := NewListByArray([]int{3, 2, 0})
	list.Next.Next.Next = list
	fmt.Println(hasCycle(list))
}
