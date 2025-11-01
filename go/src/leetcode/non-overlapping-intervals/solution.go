// 435. 无重叠区间
// https://leetcode.cn/problems/non-overlapping-intervals

package non_overlapping_intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	cnt := 1
	k := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[k][1] {
			k = i
			cnt++
		}
	}
	return len(intervals) - cnt
}
func EraseOverlapIntervals(intervals [][]int) int {
	return eraseOverlapIntervals(intervals)
}
