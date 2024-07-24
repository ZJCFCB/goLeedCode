package backtrack

import (
	"fmt"
	"strconv"
)

type Test93 struct{}

func restoreIpAddresses(s string) []string {
	var ans []string
	var temp []string
	var length int = len(s)
	var dfs func(start int)
	dfs = func(start int) {
		if len(temp) > 4 || start > length {
			return
		}
		if len(temp) == 4 && start == length {
			str := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
			ans = append(ans, str)
		}
		i := start
		for j := 1; j < 4; j++ {
			if i+j > length {
				break
			}
			tmepstr := s[i : i+j]
			if len(tmepstr) > 1 && tmepstr[0] == '0' { // 首部不能为0
				continue
			}
			if n, _ := strconv.Atoi(tmepstr); n <= 255 {
				temp = append(temp, tmepstr)
				dfs(i + j)
				temp = temp[:len(temp)-1]
			} else {
				break
			}
		}

	}
	dfs(0)
	return ans
}

func (T Test93) Run() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
