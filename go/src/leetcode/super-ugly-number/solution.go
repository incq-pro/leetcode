// 313. 超级丑数
// https://leetcode.cn/problems/super-ugly-number

package super_ugly_number

import (
	"container/heap"
	"fmt"
)

type hp [][]int

func (x hp) Len() int            { return len(x) }
func (x hp) Less(i, j int) bool  { return x[i][0] < x[j][0] }
func (x hp) Swap(i, j int)       { x[i], x[j] = x[j], x[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	hpp1 := make([][]int, len(primes))
	for i, p2 := range primes {
		hpp1[i] = make([]int, 2)
		hpp1[i][0] = p2
		hpp1[i][1] = 0
	}
	h := hp(hpp1)
	heap.Init(&h)
	for i := 1; i < n; i++ {
		m := heap.Pop(&h).([]int)
		dp[i] = m[0]
		m[0] = (m[0] / dp[m[1]]) * dp[m[1]+1]
		m[1]++
		heap.Push(&h, m)

		for {
			m = heap.Pop(&h).([]int)
			if m[0] <= dp[i] {
				m[0] = (m[0] / dp[m[1]]) * dp[m[1]+1]
				m[1]++
				heap.Push(&h, m)
			} else {
				heap.Push(&h, m)
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func NthSuperUglyNumber(n int, primes []int) int {
	return nthSuperUglyNumber(n, primes)
}
