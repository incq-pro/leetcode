// 3027. 人员站位的方案数 II
// https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-ii/

package find_the_number_of_ways_to_place_people_ii

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	n := len(points)
	p := 0
	for i := 0; i < n-1; i++ {
		ay := points[i][1]
		maxY := math.MinInt
		for j := i + 1; j < n; j++ {
			by := points[j][1]
			if by <= ay {
				if maxY < by {
					p++
					maxY = by
				}
			}
		}
	}
	return p
}

func NumberOfPairs(points [][]int) int {
	return numberOfPairs(points)
}
