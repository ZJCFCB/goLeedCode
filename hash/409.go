package hash

type Test409 struct{}

func longestPalindrome(s string) int {
	var count map[byte]int = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	var ans int
	var flag bool
	for _, value := range count {
		if value%2 == 1 {
			flag = true
		}

		ans += (value / 2) * 2
	}
	if flag {
		ans += 1
	}
	return ans
}

func (T Test409) Run() {

}
