package hot100

import "fmt"

type Test49 struct{}

func getNew(str string) string {
	var count []int = make([]int, 26)
	for _, v := range str {
		count[v-'a']++
	}

	var str1 string
	for i := 0; i < 26; i++ {
		for j := 0; j < count[i]; j++ {
			str1 += string(i + 'a')
		}
	}
	return str1
}

func groupAnagrams(strs []string) [][]string {
	var strType []int = make([]int, len(strs))
	var idx int = 0
	var count map[string]int = make(map[string]int)

	for k, v := range strs {
		temp := getNew(v)
		t, ok := count[temp]
		if ok {
			strType[k] = t
		} else {
			count[temp] = idx
			strType[k] = idx
			idx++
		}
	}
	var ans [][]string = make([][]string, idx)
	for i := 0; i < len(strs); i++ {
		ans[strType[i]] = append(ans[strType[i]], strs[i])
	}

	return ans
}

func (T Test49) Run() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
