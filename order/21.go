package order

import (
	"fmt"
)

type Test21 struct{}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans *ListNode = &ListNode{Next: nil}
	var work *ListNode = ans
	var first, second *ListNode = list1, list2
	for first != nil && second != nil {
		if first.Val < second.Val {
			work.Next = first
			first = first.Next
			work = work.Next
			work.Next = nil
		} else {
			work.Next = second
			second = second.Next
			work = work.Next
			work.Next = nil
		}
	}
	if first == nil {
		work.Next = second
	}
	if second == nil {
		work.Next = first
	}
	return ans.Next
}

func (T Test21) Run() {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	ls := mergeTwoLists(NewListByArray(l1), NewListByArray(l2))
	fmt.Println(TurnListToArray(ls))
}
