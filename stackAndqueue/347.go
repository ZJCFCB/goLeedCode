package stackAndqueue

import "fmt"

type Test347 struct{}

type Point struct {
	data int
	fre  int
}

type MonStack struct {
	stack   []Point
	stackFz []Point
	len     int
}

func (M *MonStack) Push(data int, fre int) {
	if len(M.stack) == 0 { // 直接加
		M.stack = append(M.stack, Point{data, fre})
	} else { // 要加的条件是 stack还没满 || 频率比栈顶高
		if len(M.stack) < M.len || M.stack[len(M.stack)-1].fre < fre {
			//加的话，就要考虑移位置了
			//1. 移入临时栈 2. 插入 3. 移回去
			for len(M.stack) > 0 && M.stack[len(M.stack)-1].fre < fre {
				temp := M.stack[len(M.stack)-1]
				M.stack = M.stack[:len(M.stack)-1]
				M.stackFz = append(M.stackFz, temp)
			}

			M.stack = append(M.stack, Point{data, fre})

			for len(M.stackFz) > 0 {
				temp := M.stackFz[len(M.stackFz)-1]
				M.stackFz = M.stackFz[:len(M.stackFz)-1]
				M.stack = append(M.stack, temp)
			}
			if len(M.stack) > M.len {
				M.stack = M.stack[:M.len]
			}
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	var freCount map[int]int = make(map[int]int, len(nums))
	for _, v := range nums {
		freCount[v]++
	}
	var mk MonStack = MonStack{len: k}
	for k, v := range freCount {
		mk.Push(k, v)
	}
	var result []int
	for _, v := range mk.stack {
		result = append(result, v.data)
	}
	return result
}

func (T Test347) Run() {
	nums := []int{1, 3, 3, 1, 3}
	fmt.Println(topKFrequent(nums, 2))
}
