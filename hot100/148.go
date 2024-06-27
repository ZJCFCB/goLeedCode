package hot100

import (
	"fmt"
)

type Test148 struct{}

func merge(l1, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	work1, work2, work3 := newHead, l1, l2
	for work2 != nil && work3 != nil {
		if work2.Val < work3.Val {
			work1.Next = work2
			work2 = work2.Next
		} else {
			work1.Next = work3
			work3 = work3.Next
		}
		work1 = work1.Next
	}
	if work2 == nil {
		work1.Next = work3
	} else {
		work1.Next = work2
	}
	return newHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func (test Test148) Run() {
	list := NewListByArray([]int{4, 3, 1, 2})
	x := sortList(list)
	fmt.Println(TurnListToArray(x))
}
