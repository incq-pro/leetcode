// 3074. 重新分装苹果
// https://leetcode.cn/problems/apple-redistribution-into-boxes

package apple_redistribution_into_boxes

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)
	totalApple := 0
	for _, a := range apple {
		totalApple += a
	}
	ans := 0
	m := len(capacity)
	totalCapacity := 0
	for ; ans < m && totalCapacity < totalApple; ans++ {
		totalCapacity += capacity[m-1-ans]
	}
	return ans
}

func MinimumBoxes(apple []int, capacity []int) int {
	return minimumBoxes(apple, capacity)
}
