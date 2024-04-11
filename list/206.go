package list

import (
	"fmt"
)

type Test206 struct{}

func reverseList(head *ListNode) *ListNode {
	var behind, forhead, temp *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	behind = head
	forhead = head.Next

	for forhead != nil {
		temp = forhead.Next
		forhead.Next = behind
		behind = forhead
		forhead = temp
	}
	head.Next = nil
	return behind
}

/*
func reverseList(head *ListNode) *ListNode {
    var newList *ListNode

    for head != nil{
        tmp := head.Next
        head.Next = newList
        newList = head
        head = tmp
    }
    return newList
}
*/

func (test Test206) Run() {
	array := []int{}
	head := NewListByArray(array)
	result := reverseList(head)
	ans := TurnListToArray(result)
	fmt.Println(ans)
}
