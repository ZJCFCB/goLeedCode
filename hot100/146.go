package hot100

import "fmt"

type Test146 struct {
}

type LNode struct {
	key   int
	value int
	next  *LNode
	pre   *LNode
}

type LRUCache struct {
	cache map[int]*LNode
	cap   int
	count int
	head  *LNode
	tail  *LNode
}

func Constructor(capacity int) LRUCache {
	head := &LNode{}
	tail := &LNode{}
	head.next = tail
	tail.pre = head
	return LRUCache{cap: capacity, head: head, tail: tail, cache: make(map[int]*LNode)}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}
	v.pre.next = v.next
	v.next.pre = v.pre
	v.next = this.head.next
	this.head.next.pre = v
	v.pre = this.head
	this.head.next = v
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cache[key] //查到了
	if ok {
		v.value = value
		v.pre.next = v.next
		v.next.pre = v.pre
		v.next = this.head.next
		this.head.next.pre = v
		v.pre = this.head
		this.head.next = v
	} else {
		if this.count >= this.cap {
			//先看是否需要淘汰
			temp := this.tail.pre.pre
			delete(this.cache, temp.next.key)
			temp.next = temp.next.next
			this.tail.pre = temp
			this.count--
		}
		//然后再插入
		this.count++
		newNode := &LNode{key: key, value: value}
		newNode.next = this.head.next
		this.head.next.pre = newNode
		newNode.pre = this.head
		this.head.next = newNode
		this.cache[key] = newNode
	}

}

func (T Test146) Run() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(2, 9)
	obj.Put(3, 3)
	param_1 := obj.Get(1)
	fmt.Println(param_1)
}
