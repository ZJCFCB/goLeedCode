package greedy

import "fmt"

type Test135 struct{}

// 分成左比较和右比较
func candy(ratings []int) int {
	var length = len(ratings)
	var left []int = make([]int, length)
	var right []int = make([]int, length)
	var result int = 0
	for k, v := range ratings {
		if k > 0 && v > ratings[k-1] {
			left[k] = left[k-1] + 1
		} else {
			left[k] = 1
		}
	}

	for k := length - 1; k >= 0; k-- {
		if k < length-1 && ratings[k] > ratings[k+1] {
			right[k] = right[k+1] + 1
		} else {
			right[k] = 1
		}
		result += max(right[k], left[k])
	}
	return result
}

func (T Test135) Run() {
	var ratings []int = []int{1, 3, 2, 2, 1}
	fmt.Println(candy(ratings))
}
