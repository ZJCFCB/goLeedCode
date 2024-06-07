package order

import "fmt"

type Test2 struct{}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode = &ListNode{}
	work := ans
	var work1, work2 *ListNode
	var flag int
	for work1, work2 = l1, l2; work1 != nil && work2 != nil; work1, work2 = work1.Next, work2.Next {
		temp := &ListNode{Next: nil}
		sum := work1.Val + work2.Val + flag
		flag = (sum) / 10
		temp.Val = sum % 10
		work.Next = temp
		work = work.Next
	}

	if work1 == nil {
		for w := work2; w != nil; w = w.Next {
			temp := &ListNode{Next: nil}
			temp.Val = (flag + w.Val) % 10
			flag = (flag + w.Val) / 10
			work.Next = temp
			work = work.Next
		}
	}
	if work2 == nil {
		for w := work1; w != nil; w = w.Next {
			temp := &ListNode{Next: nil}
			temp.Val = (flag + w.Val) % 10
			flag = (flag + w.Val) / 10
			work.Next = temp
			work = work.Next
		}
	}

	if flag == 1 {
		temp := &ListNode{Next: nil, Val: 1}
		work.Next = temp
	}
	return ans.Next
}

func (T Test2) Run() {
	l1 := NewListByArray([]int{0})
	l2 := NewListByArray([]int{0})
	ans := addTwoNumbers(l1, l2)
	fmt.Println(TurnListToArray(ans))
}
