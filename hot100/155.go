package hot100

import (
	"fmt"
	"math"
)

type Test155 struct{}

type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor1() MinStack {
	var min MinStack
	min.minstack = append(min.minstack, math.MaxInt)
	return min
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	} else {
		this.minstack = append(this.minstack, this.minstack[len(this.minstack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}

func (T Test155) Run() {
	obj := Constructor1()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3, param_4)

}
