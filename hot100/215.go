package hot100

import "fmt"

type Test215 struct{}

func findKthLargest(nums []int, k int) int {
	var heap []int = nums //大根堆
	var down func(int)
	down = func(n int) {
		var flag int = n
		for flag*2+1 < len(heap) {
			var next int
			next = flag*2 + 1
			if next+1 < len(heap) && heap[next+1] > heap[next] {
				next = next + 1
			}
			if heap[flag] > heap[next] {
				break
			}
			heap[next], heap[flag] = heap[flag], heap[next]
			flag = next
		}
	}
	for i := len(nums) >> 1; i >= 0; i-- {
		down(i)
	}
	for j := 0; j < k-1; j++ {
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		down(0)
	}
	return heap[0]
}

func (T Test215) Run() {
	fmt.Println(findKthLargest([]int{5, 2, 4, 1, 3, 6, 0}, 2))
}
