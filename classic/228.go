package classic

import (
	"fmt"
	"strconv"
)

type Test228 struct{}

func summaryRanges(nums []int) []string {

	var ans []string
	if len(nums) == 0 {
		return ans
	}
	var left, right int = nums[0], nums[0]
	var str string

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			right++
		} else {
			if left == right {
				str = strconv.Itoa(left)
				ans = append(ans, str)
			} else {
				str = strconv.Itoa(left) + "->" + strconv.Itoa(right)
				ans = append(ans, str)
			}
			left = nums[i]
			right = nums[i]
		}
	}
	if left == right {
		str = strconv.Itoa(left)
		ans = append(ans, str)
	} else {
		str = strconv.Itoa(left) + "->" + strconv.Itoa(right)
		ans = append(ans, str)
	}
	return ans
}

func (T Test228) Run() {
	fmt.Println(summaryRanges([]int{0}))
}
