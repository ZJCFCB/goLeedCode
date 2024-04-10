package list

import (
	"fmt"
)

type Test707 struct{}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  0,
		Next: nil,
	}
}

func NewMyLinkedList(val int) *MyLinkedList {
	return &MyLinkedList{
		Val:  val,
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	var work *MyLinkedList = this.Next
	for work != nil && index > 0 {
		work = work.Next
		index--
	}
	if index == 0 && work != nil {
		return work.Val
	} else {
		return -1
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	temp := NewMyLinkedList(val)
	temp.Next = this.Next
	this.Next = temp
}

func (this *MyLinkedList) AddAtTail(val int) {
	var work *MyLinkedList = this
	for work.Next != nil {
		work = work.Next
	}
	work.Next = NewMyLinkedList(val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	var work *MyLinkedList = this
	for work.Next != nil && index > 0 {
		work = work.Next
		index--
	}
	if index == 0 {
		temp := NewMyLinkedList(val)
		temp.Next = work.Next
		work.Next = temp
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	var work *MyLinkedList = this
	for work.Next != nil && index > 0 {
		work = work.Next
		index--
	}
	if index == 0 && work.Next != nil {
		work.Next = work.Next.Next
	}
}

func (test Test707) Run() {
	testList := Constructor()
	testList.AddAtHead(4)
	testList.Get(1)
	testList.AddAtHead(1)
	testList.AddAtHead(5)
	testList.DeleteAtIndex(3)
	testList.AddAtHead(7)
	testList.Get(3)
	testList.Get(3)
	testList.Get(3)
	testList.AddAtHead(1)
	testList.DeleteAtIndex(4)
	fmt.Println(TurnListToArrayMyLinkedList(testList.Next))

}
