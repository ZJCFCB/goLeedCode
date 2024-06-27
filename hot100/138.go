package hot100

import "fmt"

type Test138 struct{}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var point map[*Node]int = make(map[*Node]int)
	var count int
	//先完成链表的复制
	for f := head; f != nil; f = f.Next {
		point[f] = count // 记录当前point的index
		count++          //统计元素的数量
	}

	var newHead []*Node = make([]*Node, count)
	var flag int
	var work *Node = head

	//copy一遍非rodm
	for work != nil {
		newHead[flag] = &Node{Val: work.Val, Next: nil, Random: nil}
		flag++
		work = work.Next
	}

	for i := 0; i < count-1; i++ {
		newHead[i].Next = newHead[i+1]
	}

	work = head
	flag = 0
	for work != nil {
		if work.Random == nil {
			newHead[flag].Random = nil
		} else {
			rpost := point[work.Random]
			newHead[flag].Random = newHead[rpost]
		}
		flag++
		work = work.Next
	}
	return newHead[0]
}

func (T Test138) Run() {
	head := &Node{Val: 2, Next: nil}
	n := &Node{Val: 3, Next: nil}
	m := &Node{Val: 4, Next: nil}
	head.Next = n
	n.Next = m
	head.Random = m
	ans := copyRandomList(head)
	fmt.Println(ans.Next.Val)
}
