package order

import "fmt"

type Test6 struct{}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ans [][]byte = make([][]byte, numRows)
	var post int = 0
	for i := 0; i < len(s); {
		if post == 0 {
			for post < numRows && i < len(s) {
				ans[post] = append(ans[post], s[i])
				post++
				i++
			}
		}
		if post == numRows {
			post -= 2
			for post > 0 && i < len(s) {
				ans[post] = append(ans[post], s[i])
				post--
				i++
			}

		}

	}
	var result string
	for _, v := range ans {
		result += string(v)
	}
	return result
}

func (T Test6) Run() {
	s := "A"
	fmt.Println(convert(s, 4))
}
