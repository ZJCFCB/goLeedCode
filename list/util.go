package list

import (
	_ "fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 根据传入的数组，构造一个不带头节点的链表，并返回头节点
// 其中头节点的val值记录有效节点的数量
func NewListByArray(array []int) *ListNode {
	var head, work *ListNode
	head = NewListNode(0)
	work = head
	for _, value := range array {
		temp := NewListNode(value)
		work.Next = temp
		work = work.Next
	}
	return head.Next
}

// 传入一个链表的头节点，返回链表中的val以切片的形式返回
func TurnListToArray(head *ListNode) []int {
	var work *ListNode = head
	var ans []int
	for work != nil {

		ans = append(ans, work.Val)
		work = work.Next
	}
	return ans
}
