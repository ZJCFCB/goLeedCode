package order

import (
	"fmt"
)

type Test13 struct{}
type valueWeight struct {
	value  int
	weight int
}

var symbolvalue = map[byte]valueWeight{
	'I': {1, 1},
	'V': {5, 2},
	'X': {10, 3},
	'L': {50, 4},
	'C': {100, 5},
	'D': {500, 6},
	'M': {1000, 7},
}

func romanToInt(s string) int {
	var result int

	for i := 0; i < len(s); i++ {
		symbol := s[i]
		value := symbolvalue[symbol].value
		sign := 1
		if i+1 < len(s) && symbolvalue[symbol].weight < symbolvalue[s[i+1]].weight {
			sign = -1
		}
		result += sign * value
	}
	return result
}

func (T Test13) Run() {

	fmt.Println(romanToInt("MCMXCIV"))
}
