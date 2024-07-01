package hot100

import "fmt"

type Test118 struct {
}

func generate(numRows int) [][]int {
	var ans [][]int
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if numRows == 2 {
		return ans
	}
	for i := 3; i <= numRows; i++ {
		temp := []int{1}
		work := ans[len(ans)-1]
		for j := 0; j < len(work)-1; j++ {
			temp = append(temp, work[j]+work[j+1])
		}
		temp = append(temp, 1)
		ans = append(ans, temp)
	}
	return ans
}
func (T Test118) Run() {
	fmt.Println(generate(5))
}
