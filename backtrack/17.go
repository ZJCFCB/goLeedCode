package backtrack

import "fmt"

type Test17 struct{}

var numberToLetter map[byte][]string = map[byte][]string{
	'1': []string{},
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var result []string = make([]string, 0)

	var fs func(d string, flag int)

	fs = func(d string, flag int) {
		if flag == len(d) {
			return
		} else {
			var temp []string = make([]string, len(result))
			copy(temp, result)
			result = result[:0]
			work := numberToLetter[digits[flag]]
			for i := 0; i < len(work); i++ {
				if flag == 0 {
					result = append(result, work[i])
				} else {
					for _, v := range temp {
						result = append(result, v+work[i])
					}
				}
			}
			fs(digits, flag+1)
		}
	}
	fs(digits, 0)
	return result
}

func (T Test17) Run() {
	digits := "23"

	fmt.Println(letterCombinations(digits))
}
