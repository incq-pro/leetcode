// 2048. 下一个更大的数值平衡数
// https://leetcode.cn/problems/next-greater-numerically-balanced-number

package next_greater_numerically_balanced_number

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if isBeautifulNumber(i) {
			return i
		}
	}
}

func isBeautifulNumber(num int) bool {
	m := make([]int, 10)
	for num > 0 {
		m[num%10] += 1
		num = num / 10
	}
	if m[0] != 0 {
		return false
	} else {
		for k := 0; k < len(m); k++ {
			if m[k] > 0 && m[k] != k {
				return false
			}
		}
	}
	return true
}

func NextBeautifulNumber(n int) int {
	return nextBeautifulNumber(n)
}
