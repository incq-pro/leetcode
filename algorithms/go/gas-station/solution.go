// 134. 加油站
// https://leetcode.cn/problems/gas-station/

package gas_station

func mod(x, n int) int {
	if x < n {
		return x
	}
	return x % n
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	s, e := 0, 0
	left := 0
	for {
		for {
			left += gas[e] - cost[e]
			e = mod(e+1, n)
			if left < 0 {
				break
			}
			if e == s {
				return s
			}
		}
		for left < 0 {
			left -= gas[s] - cost[s]
			s = mod(s+1, n)
			if s == 0 {
				return -1
			}
			if s == e {
				break
			}
		}
	}
	return -1
}

func CanCompleteCircuit(gas []int, cost []int) int {
	return canCompleteCircuit(gas, cost)
}
