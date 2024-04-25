package hash

import "fmt"

type Test242 struct{}

func isAnagram(s string, t string) bool {
	var Hash [26]int
	for _, v := range s {
		Hash[v-'a']++
	}
	for _, v := range t {
		Hash[v-'a']--
	}
	for _, v := range Hash {
		if v != 0 {
			return false
		}
	}
	return true
}

func (Test242) Run() {
	var s, t string = "abc", "caa"
	result := isAnagram(s, t)
	fmt.Println(result)
}
