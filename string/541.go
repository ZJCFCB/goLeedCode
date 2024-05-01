package string

import "fmt"

type Test541 struct{}

// 注意一种案例： “abcdefgh”  k=20

func reverseStr(s string, k int) string {
	var result []byte = []byte(s)
	var first, second int = 0, min(k-1, len(s)-1)
	for first < len(s) {
		//开始转换
		for i, j := first, second; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
		first = second + k + 1
		second = min(first+k-1, len(s)-1)

	}
	return string(result)
}

func (T Test541) Run() {
	var s string = "abcdefgh"
	fmt.Println(reverseStr(s, 3))
}
