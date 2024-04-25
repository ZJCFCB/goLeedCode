package list

type Test142 struct{}

func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode = head, head
	var isCycle bool = false

	for {
		if slow == nil || fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
		fast = fast.Next
		if fast == slow {
			isCycle = true
			break
		}
	}
	if isCycle { // 出现了相等的情况
		fast = head
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
		return fast
	} else {
		return nil
	}
}

func (T Test142) Run() {}
