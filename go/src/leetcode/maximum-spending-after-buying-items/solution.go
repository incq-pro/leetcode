// 2931. 购买物品的最大开销
// https://leetcode.cn/problems/maximum-spending-after-buying-items/

package maximum_spending_after_buying_items

import "container/heap"

type Item struct {
	R, C, V int
}

type Q []*Item

func (q Q) Len() int {
	return len(q)
}

func (q Q) Less(i, j int) bool {
	return q[i].V < q[j].V
}

func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Q) Push(x any) {
	*q = append(*q, x.(*Item))
}

func (q *Q) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	q := &Q{}
	heap.Init(q)
	for i := 0; i < m; i++ {
		heap.Push(q, &Item{i, n - 1, values[i][n-1]})
	}
	total := m * n
	var cost int64
	for i := 1; i <= total; i++ {
		x := heap.Pop(q).(*Item)
		cost += int64(x.V * i)
		if x.C > 0 {
			x.C--
			x.V = values[x.R][x.C]
			heap.Push(q, x)
		}
	}
	return cost
}

func MaxSpending(values [][]int) int64 {
	return maxSpending(values)
}
