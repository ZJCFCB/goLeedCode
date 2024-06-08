package order

import "fmt"

type Test3 struct{}

func lengthOfLongestSubstring(s string) int {
	var count map[byte]int = make(map[byte]int)
	var ans int

	for i := 0; i < len(s); i++ {
		fmt.Println(count)
		post, ok := count[s[i]]

		if ok { // 重复
			count = make(map[byte]int)
		} else {
			count[s[i]] = i
		}

		ans = max(ans, len(count))
		if ok {
			i = post
		}
	}
	return ans
}

func (T Test3) Run() {
	ans := "abacd"
	fmt.Println(lengthOfLongestSubstring(ans))
}
