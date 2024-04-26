package stackAndqueue

import "fmt"

type Test20 struct {
}

func isValid(s string) bool {
	var stack []rune

	for _, v := range s {
		if len(stack) == 0 || v == '(' || v == '[' || v == '{' { //入栈
			stack = append(stack, v)
		} else {
			switch v {
			case ')':
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

//击败100%的代码
//确实非常优雅，妙
/*

func isValid(s string) bool {
	var tmp []rune
	a := make(map[rune]rune)
	a[')'] = '('
	a[']'] = '['
	a['}'] = '{'
	if len(s) == 0 {
		return true
	} else if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}
	for _, m := range s {
		if m == '(' || m == '[' || m == '{' {
			tmp = append(tmp, m)
		} else if len(tmp) > 0 && a[m] == tmp[len(tmp)-1] {
			tmp = tmp[:len(tmp)-1]
		} else {
			return false
		}
	}
	if len(tmp) == 0 {
		return true
	}
	return false
}

*/

func (T Test20) Run() {
	var s string = "()([]{}"

	fmt.Println(isValid(s))
}
