package list

type Test160 struct{}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var flagA, flagB *ListNode = headA, headB
	for flagA != nil && flagB != nil {
		flagA = flagA.Next
		flagB = flagB.Next
	}
	if flagA == nil {
		flagA = headB
		for flagB != nil {
			flagA = flagA.Next
			flagB = flagB.Next
		}
		flagB = headA
	} else if flagB == nil {
		flagB = headA
		for flagA != nil {
			flagA = flagA.Next
			flagB = flagB.Next
		}
		flagA = headB
	}

	for flagA != nil && flagA != flagB {
		flagA = flagA.Next
		flagB = flagB.Next
	}
	return flagA
}

func (T Test160) Run() {
}
