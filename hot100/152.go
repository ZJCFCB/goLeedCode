package hot100

import (
	"fmt"
)

type Test152 struct{}

func maxProduct(nums []int) int {
	var ans, tempMax, tempMin int = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmax, tmin := tempMax, tempMin
		tempMax = max(nums[i]*tmax, nums[i], tmin*nums[i])
		tempMin = min(nums[i]*tmin, nums[i], tmax*nums[i])
		ans = max(ans, tempMax)
		fmt.Println(tempMax, tempMin, ans)
	}
	return ans
}

func (T Test152) Run() {
	fmt.Println(maxProduct([]int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}))
}
