package hash

import (
	"fmt"
	"sort"
)

type Test18 struct{}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	var result [][]int = make([][]int, 0)

	for i := 0; i < length-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if (i > 0 && nums[i] == nums[i-1]) || (nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target) {
			continue
		}

		for j := i + 1; j < length-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if (j > i+1 && nums[j] == nums[j-1]) || nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			for left, right := j+1, length-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ { //相同的就排除掉吧

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- { //相同的就排除掉吧

					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func (T Test18) Run() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(fourSum(nums, -1))
}
