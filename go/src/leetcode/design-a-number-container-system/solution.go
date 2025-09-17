// 2349. 设计数字容器系统
// https://leetcode.cn/problems/design-a-number-container-system

package design_a_number_container_system

import "container/heap"

type hp []int

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

type NumberContainers struct {
	index  map[int]int
	number map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{
		index:  map[int]int{},
		number: map[int]*hp{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if v, ok := this.index[index]; ok {
		if v == number {
			return
		}
	}
	this.index[index] = number
	if v, ok := this.number[number]; ok {
		heap.Push(v, index)
	} else {
		h := hp([]int{})
		v = &h
		heap.Init(v)
		this.number[number] = v
		heap.Push(v, index)
	}
}

func (this *NumberContainers) Find(number int) int {
	if v, ok := this.number[number]; !ok {
		return -1
	} else {
		for v.Len() > 0 {
			i := heap.Pop(v).(int)
			if v2, ok2 := this.index[i]; ok2 && v2 == number {
				heap.Push(v, i)
				return i
			}
		}
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
