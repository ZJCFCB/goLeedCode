package hot100

import (
	"fmt"
)

type Test295 struct{}

type MedianFinder struct {
	minheap, maxheap []int
}

func (m MedianFinder) minheapUp() {
	var flag int = len(m.minheap) - 1
	var temp int
	for flag >= 0 {
		temp = (flag - 1) >> 1
		if temp >= 0 && m.minheap[temp] > m.minheap[flag] {
			m.minheap[temp], m.minheap[flag] = m.minheap[flag], m.minheap[temp]
			flag = temp
		} else {
			break
		}
	}

}

func (m MedianFinder) minheapDown() {
	var flag int
	var next int
	for 2*flag+1 < len(m.minheap) {
		next = 2*flag + 1
		if next+1 < len(m.minheap) && m.minheap[next+1] < m.minheap[next] {
			next = next + 1
		}
		if m.minheap[flag] > m.minheap[next] {
			m.minheap[flag], m.minheap[next] = m.minheap[next], m.minheap[flag]
			flag = next
		} else {
			break
		}
	}
}

func (m MedianFinder) maxheapUp() {
	var flag int = len(m.maxheap) - 1
	var temp int
	for flag >= 0 {
		temp = (flag - 1) >> 1
		if temp >= 0 && m.maxheap[temp] < m.maxheap[flag] {
			m.maxheap[temp], m.maxheap[flag] = m.maxheap[flag], m.maxheap[temp]
			flag = temp
		} else {
			break
		}
	}
}

func (m MedianFinder) maxheapDown() {
	var flag int
	var next int
	for 2*flag+1 < len(m.maxheap) {
		next = 2*flag + 1
		if next+1 < len(m.maxheap) && m.maxheap[next+1] > m.maxheap[next] {
			next = next + 1
		}
		if m.maxheap[flag] < m.maxheap[next] {
			m.maxheap[flag], m.maxheap[next] = m.maxheap[next], m.maxheap[flag]
			flag = next
		} else {
			break
		}
	}

}

func Constructor5() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxheap) == 0 {
		this.maxheap = append(this.maxheap, num)
		this.maxheapUp()
	} else {
		if num <= this.maxheap[0] {
			this.maxheap = append(this.maxheap, num)
			this.maxheapUp()
			if len(this.maxheap)-1 > len(this.minheap) {
				temp := this.maxheap[0]
				this.minheap = append(this.minheap, temp)
				this.minheapUp()
				this.maxheap[0] = this.maxheap[len(this.maxheap)-1]
				this.maxheap = this.maxheap[:len(this.maxheap)-1]
				this.maxheapDown()
			}
		} else {
			this.minheap = append(this.minheap, num)
			this.minheapUp()
			if len(this.minheap) > len(this.maxheap) {
				temp := this.minheap[0]
				this.maxheap = append(this.maxheap, temp)
				this.maxheapUp()
				this.minheap[0] = this.minheap[len(this.minheap)-1]
				this.minheap = this.minheap[:len(this.minheap)-1]
				this.minheapDown()
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxheap) == len(this.minheap) { // 偶数
		return (float64(this.minheap[0]) + float64(this.maxheap[0])) / 2
	} else {
		return float64(this.maxheap[0])
	}
}

func (this *MedianFinder) printlnheap() {
	fmt.Println(this.minheap, this.maxheap)
}

func (T Test295) Run() {
	obj := Constructor5()
	temp := []int{40, 12, 16, 14, 35, 19, 34, 35, 28, 35, 26, 6, 8, 2, 14, 25, 25, 4, 33, 18, 10, 14, 27, 3, 35, 13, 24, 27, 14, 5, 0, 38, 19, 25, 11, 14, 31, 30, 11, 31, 0}
	for _, v := range temp {
		obj.AddNum(v)
		obj.printlnheap()
		param_2 := obj.FindMedian()
		fmt.Println(param_2)

	}

}
