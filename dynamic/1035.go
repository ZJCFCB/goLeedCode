package dynamic

import "fmt"

type Test1035 struct{}

func maxUncrossedLines(nums1 []int, nums2 []int) int {

	var dp [][]int = make([][]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		dp[i] = make([]int, len(nums1))
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			for j := i; j < len(nums1); j++ {
				dp[0][j] = 1
			}
			break
		}
	}

	for i := 0; i < len(nums2); i++ {
		if nums1[0] == nums2[i] {
			for j := i; j < len(nums2); j++ {
				dp[j][0] = 1
			}
			break
		}
	}

	for i := 1; i < len(nums2); i++ {
		for j := 1; j < len(nums1); j++ {
			if nums1[j] == nums2[i] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[j-1][i])
			}
		}
	}
	return dp[len(nums2)-1][len(nums1)-1]
}

func (T Test1035) Run() {
	nums1 := []int{1, 4, 2, 4}
	nums2 := []int{1, 4, 2, 4}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
