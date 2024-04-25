package hash

import "fmt"

type Test1002 struct {
}

func commonChars(words []string) []string {
	var result [26]int
	for _, v := range words[0] {
		result[v-'a']++
	}

	for i := 1; i < len(words); i++ {
		var temp [26]int
		for _, v := range words[i] {
			temp[v-'a']++
		}

		for j := 0; j < 26; j++ {
			if result[j] > temp[j] {
				result[j] = temp[j]
			}
		}
	}
	var answer []string
	for i := 0; i < 26; i++ {
		if result[i] != 0 {
			for j := result[i]; j > 0; j-- {
				answer = append(answer, string('a'+i))
			}
		}
	}
	return answer
}

func (T *Test1002) Run() {
	var t []string = []string{"cool", "lock", "cook"}
	r := commonChars(t)
	fmt.Println(r)

}
