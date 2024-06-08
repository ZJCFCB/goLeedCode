package order

import "fmt"

type Test4 struct{}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merge []int = make([]int, len(nums1)+len(nums2))
	var f1, f2, fm int
	for f1 < len(nums1) && f2 < len(nums2) {
		if nums1[f1] > nums2[f2] {
			merge[fm] = nums2[f2]
			fm++
			f2++
		} else {
			merge[fm] = nums1[f1]
			fm++
			f1++
		}
	}
	if f1 == len(nums1) {
		for f2 < len(nums2) {
			merge[fm] = nums2[f2]
			fm++
			f2++
		}
	}
	if f2 == len(nums2) {
		for f1 < len(nums1) {
			merge[fm] = nums1[f1]
			fm++
			f1++
		}
	}
	length := len(merge)
	if length%2 == 1 {
		return float64(merge[length/2])
	} else {
		return (float64(merge[length/2]) + float64(merge[length/2-1])) / 2
	}
}

func (T Test4) Run() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
