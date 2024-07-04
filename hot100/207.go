package hot100

import "fmt"

type Test207 struct{}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var count map[int][]int = make(map[int][]int, len(prerequisites))
	var rd map[int]int = make(map[int]int)
	for i := 0; i < numCourses; i++ {
		rd[i] = 0
	}
	for i := 0; i < len(prerequisites); i++ {
		count[prerequisites[i][1]] = append(count[prerequisites[i][1]], prerequisites[i][0])
		rd[prerequisites[i][0]]++
	}

	var flag bool = true
	for flag {
		flag = false
		for k, v := range rd {
			if v == 0 {
				flag = true //代表还有入度为0的节点
				for _, value := range count[k] {
					rd[value]--
				}
				delete(rd, k)
			}
		}

	}
	if len(rd) == 0 {
		return true
	}
	return false
}

func (T Test207) Run() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
}
