// 7. 整数反转
// https://leetcode.cn/problems/reverse-integer/

package reverse_integer

import "math"

func reverse(x int) int {
	result := 0
	for x != 0 {
		remainder := x % 10
		x = x / 10
		if result > 0 && result > (math.MaxInt32-remainder)/10 {
			return 0
		}
		if result < 0 && result < (math.MinInt32-remainder)/10 {
			return 0
		}
		result = result*10 + remainder
	}
	return result
}

func Reverse(x int) int {
	return reverse(x)
}
