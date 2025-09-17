// 264. 丑数 II
// https://leetcode.cn/problems/ugly-number-ii/

package ugly_number_ii

import "math"

func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	factor := [3]int{2, 3, 5}
	for i := 1; i < n; i++ {
		m := math.MaxInt
		last := nums[i-1]
		for _, f := range factor {
			for j := i - 1; j >= 0; j-- {
				v := nums[j] * f
				if v <= last {
					break
				}
				m = min(m, v)
			}
		}
		nums[i] = m
	}
	return nums[n-1]
}

func NthUglyNumber(n int) int {
	return nthUglyNumber(n)
}
