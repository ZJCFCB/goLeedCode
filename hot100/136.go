package hot100

func singleNumber(nums []int) int {
	var ans int = 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
