package hash

import "fmt"

type Test454 struct{}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	var hashCount1 map[int]int = make(map[int]int, n*n)
	//这里注意一下写法上的细节，既然明确了hashCount1的最大长度
	//一次分配避免扩容，可以带来很多时间上的优势
	var result int = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			hashCount1[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += hashCount1[-nums3[i]-nums4[j]]
		}
	}
	return result
}

func (T Test454) Run() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
