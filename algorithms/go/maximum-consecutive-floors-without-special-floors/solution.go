//  2274. 不含特殊楼层的最大连续楼层数
// https://leetcode.cn/problems/maximum-consecutive-floors-without-special-floors/

package maximum_consecutive_floors_without_special_floors

import "slices"

func maxConsecutive(bottom int, top int, special []int) int {
	slices.Sort(special)
	m := 0
	m = max(m, special[0]-bottom)
	m = max(m, top-special[len(special)-1])
	n := len(special)
	for i := 0; i < n-1; i++ {
		m = max(m, special[i+1]-special[i]-1)
	}
	return m
}

func MaxConsecutive(bottom int, top int, special []int) int {
	return maxConsecutive(bottom, top, special)
}
