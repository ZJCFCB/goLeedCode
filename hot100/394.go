package hot100

import "fmt"

type Test394 struct {
}

func decodeString(s string) string {
	var stackstring []string
	var stacknum []byte
	var flag int
	for flag < len(s) {
		switch {
		case '0' <= s[flag] && s[flag] <= '9': // 是数字
			stacknum = append(stacknum, s[flag])
		case 'a' <= s[flag] && s[flag] <= 'z':
			stackstring = append(stackstring, string(s[flag]))
		case s[flag] == '[':
			stacknum = append(stacknum, s[flag])
			stackstring = append(stackstring, string(s[flag]))
		case s[flag] == ']':
			var substr string
			//获得当前的子串
			for len(stackstring) != 0 && stackstring[len(stackstring)-1] != "[" {
				substr = stackstring[len(stackstring)-1] + substr
				stackstring = stackstring[:len(stackstring)-1]
			}
			stackstring = stackstring[:len(stackstring)-1]
			//获得子串所重复的次数
			var num []int
			stacknum = stacknum[:len(stacknum)-1]
			for len(stacknum) != 0 && stacknum[len(stacknum)-1] != '[' {
				num = append(num, int(stacknum[len(stacknum)-1]-'0'))
				stacknum = stacknum[:len(stacknum)-1]
			}

			var number int
			for len(num) != 0 {
				number = number*10 + num[len(num)-1]
				num = num[:len(num)-1]
			}

			var ss string
			for i := 0; i < number; i++ {
				ss += substr
			}
			stackstring = append(stackstring, ss)
		}
		flag++
	}
	var ans string

	for i := 0; i < len(stackstring); i++ {
		ans += stackstring[i]
	}
	return ans
}

func (T Test394) Run() {
	str := "3[a]2[bc]"
	fmt.Println(decodeString(str))
}
