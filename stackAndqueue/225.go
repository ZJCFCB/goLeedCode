package stackAndqueue

import "fmt"

type Test225 struct{}

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor_() MyStack {
	return MyStack{queue1: make([]int, 0), queue2: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	if len(this.queue2) != 0 { //数据全部在Q2
		this.queue2 = append(this.queue2, x)
	} else {
		this.queue1 = append(this.queue1, x)
	}
}

func (this *MyStack) TurnQueue() int {
	if len(this.queue2) == 0 {
		length := len(this.queue1) - 1
		for i := 0; i < length; i++ {
			this.queue2 = append(this.queue2, this.queue1[0])
			this.queue1 = this.queue1[1:]
		}
		return 1
	} else {
		length := len(this.queue2) - 1
		for i := 0; i < length; i++ {
			this.queue1 = append(this.queue1, this.queue2[0])
			this.queue2 = this.queue2[1:]
		}
		return 2
	}
}

func (this *MyStack) Pop() int {
	flag := this.TurnQueue()
	var result int
	if flag == 1 {
		result = this.queue1[0]
		this.queue1 = this.queue1[1:]
	} else {
		result = this.queue2[0]
		this.queue2 = this.queue2[1:]
	}
	return result
}

func (this *MyStack) Top() int {
	flag := this.TurnQueue()
	var result int
	if flag == 1 {
		result = this.queue1[0]
		this.queue2 = append(this.queue2, result)
		this.queue1 = this.queue1[1:]
	} else {
		result = this.queue2[0]
		this.queue1 = append(this.queue1, result)
		this.queue2 = this.queue2[1:]
	}
	return result
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

func (T Test225) Run() {
	myqueue := Constructor_()
	myqueue.Push(1)
	myqueue.Push(2)
	fmt.Println(myqueue.Top())
	fmt.Println(myqueue.Pop())
	fmt.Println(myqueue.Empty())
	fmt.Println(myqueue.Top())
}
