// 3066. 超过阈值的最少操作数 II
// https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-ii/

package minimum_operations_to_exceed_threshold_value_ii

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations(nums []int, k int) int {
	h := &IntHeap{}
	for _, v := range nums {
		heap.Push(h, v)
	}
	count := 0
	for len(*h) >= 2 {
		v1 := heap.Pop(h).(int)
		if v1 < k {
			v2 := heap.Pop(h).(int)
			v3 := v1*2 + v2
			heap.Push(h, v3)
			count++
		} else {
			break
		}
	}
	return count
}

func MinOperations(nums []int, k int) int {
	return minOperations(nums, k)
}
