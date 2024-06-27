package hot100

import "fmt"

type Test234 struct{}

func isPalindrome(head *ListNode) bool {
	var count []int
	work := head
	for work != nil {
		count = append(count, work.Val)
		work = work.Next
	}
	var front, rear int = 0, len(count) - 1

	for front < rear {
		if count[front] != count[rear] {
			return false
		}
		front++
		rear--
	}
	return true
}
func (T Test234) Run() {
	str1 := []int{1}
	x := isPalindrome(NewListByArray(str1))
	fmt.Println(x)
}
