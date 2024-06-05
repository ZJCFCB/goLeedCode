package dynamic

import "fmt"

type Test718 struct{}

func findLength(nums1 []int, nums2 []int) int {
	var dp [][]int = make([][]int, len(nums2))

	for i := 0; i < len(nums2); i++ {
		dp[i] = make([]int, len(nums1))
	}

	var result int

	for i := 0; i < len(nums2); i++ {
		if nums2[i] == nums1[0] {
			dp[i][0] = 1
			result = 1
		}
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[0][i] = 1
			result = 1
		}
	}

	for i := 1; i < len(nums2); i++ {
		for j := 1; j < len(nums1); j++ {
			if nums1[j] == nums2[i] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
func (T Test718) Run() {
	nums := []int{1, 2, 3, 2, 8}
	nums2 := []int{5, 6, 1, 4, 7}
	fmt.Println(findLength(nums, nums2))
}
