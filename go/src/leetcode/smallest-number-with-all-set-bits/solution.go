// 3370. 仅含置位位的最小整数
// https://leetcode.cn/problems/smallest-number-with-all-set-bits/

package smallest_number_with_all_set_bits

func smallestNumber(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n
}

func SmallestNumber(n int) int {
	return smallestNumber(n)
}
