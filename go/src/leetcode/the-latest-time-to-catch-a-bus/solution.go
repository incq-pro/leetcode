// 2332. 坐上公交的最晚时间
// https://leetcode.cn/problems/the-latest-time-to-catch-a-bus/

package the_latest_time_to_catch_a_bus

import "slices"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	slices.Sort(buses)
	slices.Sort(passengers)
	// passenger 指针, 指向将要处理的passenger
	p := 0
	n := len(buses)
	m := len(passengers)
	// 容量
	c := 0
	for b := 0; b < n; b++ {
		for c = 0; c < capacity; c++ {
			if p < m && passengers[p] <= buses[b] {
				p++
			} else {
				break
			}
		}
	}
	if c >= capacity {
		// 把最后一个人请出公交车
		p--
	}
	r := buses[n-1]
	if p < m && p >= 0 && passengers[p] <= r {
		r = passengers[p] - 1
	}
	// 保证r不重复
	for x := m - 1; x >= 0; x-- {
		if passengers[x] < r {
			break
		} else if passengers[x] == r {
			r--
		}
	}
	return r
}

func LatestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	return latestTimeCatchTheBus(buses, passengers, capacity)
}
