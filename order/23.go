package order

import (
	"fmt"
)

type Test23 struct{}

type heap struct {
	Data []*ListNode
}

func (h *heap) push(v *ListNode) {
	h.Data = append(h.Data, v)
	//然后从下往上调整位置
	post := len(h.Data) - 1
	for post > 0 && h.Data[post].Val < h.Data[(post-1)/2].Val {
		h.Data[post], h.Data[(post-1)/2] = h.Data[(post-1)/2], h.Data[post]
		post = (post - 1) / 2
	}
}
func (h *heap) getTop() *ListNode {
	return h.Data[0]
}
func (h *heap) pop() *ListNode {
	v := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	//从上往下调整
	post := 0
	for post*2+1 < len(h.Data) {
		min := post*2 + 1
		if post*2+2 < len(h.Data) && h.Data[post*2+1].Val > h.Data[post*2+2].Val {
			min = post*2 + 2
		}
		if h.Data[post].Val > h.Data[min].Val {
			h.Data[post], h.Data[min] = h.Data[min], h.Data[post]
			post = min
		} else {
			break
		}
	}
	return v
}

func (h *heap) printheap() {
	for i := 0; i < len(h.Data); i++ {
		fmt.Print(h.Data[i].Val)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {

	var f func() int
	f = func() int {
		var min int

		for i := 1; i < len(lists); i++ {
			if lists[i].Val < lists[min].Val {
				min = i
			}
		}
		return min
	}
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var h heap
	for i := 0; i < len(lists); i++ {
		temp := lists[i]
		if temp == nil {
			continue
		}
		lists[i] = lists[i].Next
		temp.Next = nil
		h.push(temp)
	}

	var ans *ListNode = new(ListNode)
	var work *ListNode = ans
	for {
		//先除去为空的链表
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				lists = append(lists[:i], lists[i+1:]...)
				i--
			}
		}
		if len(lists) == 0 {
			break
		}

		min := f()

		work.Next = h.pop()
		work = work.Next

		temp := lists[min]
		lists[min] = lists[min].Next
		temp.Next = nil
		h.push(temp)
		//h.printheap()
	}

	//把堆掏空
	for len(h.Data) > 0 {
		work.Next = h.pop()
		work = work.Next
	}
	return ans.Next
}

func (T Test23) Run() {
	l1 := NewListByArray([]int{1, 4, 5})
	l2 := NewListByArray([]int{1, 3, 4, 7, 7, 7, 7, 7, 7, 7})
	l3 := NewListByArray([]int{2, 6})
	ls := mergeKLists([]*ListNode{l1, l2, l3})
	lt := mergeKLists([]*ListNode{NewListByArray([]int{}), NewListByArray([]int{})})
	fmt.Println(TurnListToArray(ls))
	fmt.Println(TurnListToArray(lt))
}
