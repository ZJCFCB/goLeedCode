package classic

import "fmt"

type Test274 struct{}

func hIndex(citations []int) int {
	var count map[int]int = make(map[int]int)
	for _, v := range citations {
		for i := 1; i <= v; i++ {
			count[i]++
		}
	}
	var h int
	for k, v := range count {
		h = max(h, min(k, v))
	}
	return h
}

func (T Test274) Run() {
	nums := []int{2, 0, 6, 1, 5}
	fmt.Println(hIndex(nums))
}
