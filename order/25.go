package order

import (
	"fmt"
)

type Test25 struct{}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversList(head *ListNode) (*ListNode, *ListNode) {
	first, second := head, head.Next
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	head.Next = nil
	return first, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead *ListNode = &ListNode{Next: head}
	var first, second *ListNode
	first = newHead
	second = newHead

	for first != nil && second != nil {
		for i := 0; i < k; i++ {
			if second.Next == nil {
				return newHead.Next
			}
			second = second.Next
		}
		fmt.Println(first.Val, second.Val)
		temp1, temp2 := first.Next, second.Next
		second.Next = nil
		h, t := reversList(temp1)
		//fmt.Println(h.Val, t.Val)
		first.Next = h
		t.Next = temp2
		first, second = t, t
	}
	return newHead.Next
}

func (T Test25) Run() {
	l1 := NewListByArray([]int{1, 2, 3, 4, 5, 6})
	// h, t := reversList(l1)
	// fmt.Println(TurnListToArray(h))
	// fmt.Println(TurnListToArray(t))
	fmt.Println(TurnListToArray(reverseKGroup(l1, 3)))
}
