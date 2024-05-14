package greedy

import "fmt"

type Test55 struct{}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var fundMax func(i, j int) int

	fundMax = func(i, j int) int {
		var flag, work int
		flag = i
		work = nums[i]
		for x := i + 1; x < len(nums) && x <= j; x++ {
			if flag+work <= x+nums[x] {
				flag = x
				work = nums[x]
			}
		}
		return flag
	}

	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			break
		}
		if i+nums[i] >= len(nums) {
			return true
		}
		i = fundMax(i, i+nums[i])
	}

	return false

}

//看看别人的优雅代码
/*

每次都记录当前能到的最远值

func canJump(nums []int) bool {
    n := len(nums)
    k := 0

    for i:=0; i < n; i++ {
        if i > k {
            return false
        }
        k = max(k, i+nums[i])
    }
    return true
}


*/
func (T Test55) Run() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
