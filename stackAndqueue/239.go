package stackAndqueue

import "fmt"

type Test239 struct{}

// 单调队列
type MonQueue struct {
	Index []int
	Data  []int
}

func (M *MonQueue) Push(x int, index int) {
	for len(M.Data) > 0 && M.Data[len(M.Data)-1] < x { // 大于零 并且队尾元素小于新插入元素
		M.Data = M.Data[:len(M.Data)-1]
		M.Index = M.Index[:len(M.Index)-1]
	}
	//插入新元素
	M.Index = append(M.Index, index)
	M.Data = append(M.Data, x)
}

func (M *MonQueue) Pop() {
	M.Data = M.Data[1:]
	M.Index = M.Index[1:]
}

func (M *MonQueue) Front() (x, index int) {
	return M.Data[0], M.Index[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var x, index int

	//先处理前k个元素
	var mq *MonQueue = &MonQueue{}
	for i := 0; i < k; i++ {
		mq.Push(nums[i], i)
	}
	x, index = mq.Front()
	result = append(result, x)

	for i := k; i < len(nums); i++ {
		x, index = mq.Front() //先检查队首元素是否还在滑动窗口内
		if i-index+1 > k {
			mq.Pop()
		}
		mq.Push(nums[i], i)
		x, index = mq.Front()
		result = append(result, x)
	}

	return result
}

func (T Test239) Run() {
	var nums []int = []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}
