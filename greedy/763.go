package greedy

import (
	"fmt"
)

type Test763 struct{}

func partitionLabels(s string) []int {
	var result []int
	var position map[byte]int = make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		position[s[i]] = i
	}
	var begin, end int = 0, 0
	for begin < len(s) {
		end = position[s[begin]]
		for j := begin + 1; j < end; j++ {
			if position[s[j]] > end { //有更长的范围
				end = position[s[j]]
				if end == len(s)-1 {
					break
				}
			}
		}
		result = append(result, end-begin+1)
		begin = end + 1
	}
	return result
}

func (T Test763) Run() {
	var s string = "eccbbbbdec"
	fmt.Println(partitionLabels(s))
}
