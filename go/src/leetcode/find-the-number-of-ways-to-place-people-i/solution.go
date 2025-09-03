// 3025. 人员站位的方案数 I
// https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-i/

package find_the_number_of_ways_to_place_people_i

import (
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
		for j := i + 1; j < n; j++ {
			by := points[j][1]
			if ay >= by {
				bx := points[j][0]
				found := false
				for k := i + 1; k < n && points[k][0] <= bx; k++ {
					if k == j {
						continue
					}
					ky := points[k][1]
					if ky >= by && ky <= ay {
						found = true
						break
					}
				}
				if !found {
					p++
				}
			}
		}
	}
	return p
}

func NumberOfPairs(points [][]int) int {
	return numberOfPairs(points)
}
