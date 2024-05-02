package hash

import "fmt"

type Test15 struct{}

func SortThree(i, j, k int) [3]int { // 三个元素从小到大排个序
	if i > j {
		j, i = i, j
	}
	if i > k {
		i, k = k, i
	}
	if j > k {
		j, k = k, j
	}
	return [3]int{i, j, k}
}

func threeSum(nums []int) [][]int {
	var hashMap map[int][]int = make(map[int][]int)

	var tempResult map[[3]int]bool = make(map[[3]int]bool)

	//先做hash，保存对应关系
	for k, v := range nums {
		hashMap[v] = append(hashMap[v], k)
	}

	var i, j int

	for i = 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j = i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			value, ok := hashMap[0-sum]
			if ok { // 找到了
				for _, v := range value {
					if v > i && v > j {
						tempResult[SortThree(nums[i], nums[j], nums[v])] = true
					}
				}
			}
		}
	}

	var result [][]int = make([][]int, 0, len(tempResult))

	for k, _ := range tempResult {
		result = append(result, k[:])
	}
	return result
}

func (T Test15) Run() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
