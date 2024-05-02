package hash

import "fmt"

type Test383 struct{}

func canConstruct(ransomNote string, magazine string) bool {
	var hashRansomNote, hashMagazine map[rune]int = make(map[rune]int, len(ransomNote)), make(map[rune]int, len(magazine))

	for _, v := range ransomNote {
		hashRansomNote[v]++
	}

	for _, v := range magazine {
		hashMagazine[v]++
	}

	for k, v := range hashRansomNote {
		value, ok := hashMagazine[k]
		if !(ok && v <= value) {
			return false
		}
	}
	return true
}

func (T Test383) Run() {
	str1 := "abc"
	str2 := "abcccccc"

	fmt.Println(canConstruct(str1, str2))

}
