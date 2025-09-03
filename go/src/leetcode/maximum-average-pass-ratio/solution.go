// 1792. 最大平均通过率
// https://leetcode.cn/problems/maximum-average-pass-ratio/

package maximum_average_pass_ratio

import "container/heap"

type ClassHeap [][]int

func (h ClassHeap) Len() int {
	return len(h)
}

func (h ClassHeap) Less(i, j int) bool {
	return h.increment(i) > h.increment(j)
}

func (h ClassHeap) increment(i int) float64 {
	return float64(h[i][0]+1)/float64(h[i][1]+1) - float64(h[i][0])/float64(h[i][1])
}

func (h ClassHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ClassHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *ClassHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	n := len(classes)
	h := ClassHeap(classes)
	heap.Init(&h)
	for ; extraStudents > 0; extraStudents-- {
		c := heap.Pop(&h).([]int)
		c[0] += 1
		c[1] += 1
		heap.Push(&h, c)
	}
	var totalRate float64
	for _, c := range classes {
		totalRate += float64(c[0]) / float64(c[1])
	}
	return totalRate / float64(n)
}

func MaxAverageRatio(classes [][]int, extraStudents int) float64 {
	return maxAverageRatio(classes, extraStudents)
}
