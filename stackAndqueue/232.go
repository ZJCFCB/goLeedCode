package stackAndqueue

import "fmt"

type Test232 struct{}

type MyQueue struct {
	stackin  []int
	stackout []int
}

func Constructor() MyQueue {
	return MyQueue{stackin: make([]int, 0), stackout: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.stackin = append(this.stackin, x)
}

func (this *MyQueue) In2Out() {
	for len(this.stackin) != 0 {
		this.stackout = append(this.stackout, this.stackin[len(this.stackin)-1])
		this.stackin = this.stackin[:len(this.stackin)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.stackout) == 0 {
		this.In2Out()
	}
	result := this.stackout[len(this.stackout)-1]
	this.stackout = this.stackout[:len(this.stackout)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.stackout) == 0 {
		this.In2Out()
	}
	return this.stackout[len(this.stackout)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stackout) == 0 && len(this.stackin) == 0 {
		return true
	} else {
		return false
	}
}

func (T *Test232) Run() {
	Myq := Constructor()
	Myq.Push(1)
	Myq.Push(2)
	fmt.Println(Myq.stackin)
	fmt.Println(Myq.Peek())
	fmt.Println(Myq.Pop())
	fmt.Println(Myq.Peek())
	fmt.Println(Myq.Empty())
}
