// 397. 整数替换
// https://leetcode.cn/problems/integer-replacement

package integer_replacement

func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if (n & 1) == 0 {
			n = n / 2
		} else {
			p := (n + 1) / 2
			if (p&1) == 0 && p < n-1 {
				n = n + 1
			} else {
				n = n - 1
			}
		}
		count++
	}
	return count
}

func IntegerReplacement(n int) int {
	return integerReplacement(n)
}
