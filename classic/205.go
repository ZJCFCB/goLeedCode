package classic

import "fmt"

type Test205 struct{}

func isIsomorphic(s string, t string) bool {
	var count map[byte]byte = make(map[byte]byte, len(s))
	var deape map[byte]bool = make(map[byte]bool)
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		value, ok := count[s[i]]
		if !ok {
			if deape[t[i]-'0'] {
				return false
			}
			count[s[i]] = t[i]
			deape[t[i]-'0'] = true
		} else {
			if t[i] != value {
				return false
			}
		}
	}
	return true
}

func (T Test205) Run() {
	fmt.Println(isIsomorphic("abb", "dee"))
}
