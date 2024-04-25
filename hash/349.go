package hash

import "fmt"

type Test349 struct{}

func intersection(nums1 []int, nums2 []int) []int {
	var ans map[int]bool = make(map[int]bool, 100)
	var result []int
	for _, v := range nums1 {
		ans[v] = true
	}

	for _, v := range nums2 {
		_, ok := ans[v]
		if ok {
			delete(ans, v)
			result = append(result, v)
		}
	}
	return result
}

func (T Test349) Run() {
	var nums1 []int = []int{4, 9, 5}
	var nums2 []int = []int{9, 4, 9, 8, 4}

	result := intersection(nums1, nums2)
	fmt.Println(result)

}
