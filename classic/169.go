package classic

import "fmt"

type Test169 struct{}

func majorityElement(nums []int) int {
	var can, count int = -1, 0

	for _, v := range nums {
		if can == v {
			count++
		} else {
			count--
			if count < 0 {
				can = v
				count = 0
			}
		}
	}
	return can
}
func (T Test169) Run() {
	nums := []int{2, 2, 1, 1, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
