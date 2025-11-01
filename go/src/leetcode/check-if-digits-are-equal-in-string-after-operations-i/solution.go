// 3461. 判断操作后字符串中的数字是否相等 I
// https://leetcode.cn/problems/check-if-digits-are-equal-in-string-after-operations-i

package check_if_digits_are_equal_in_string_after_operations_i

func hasSameDigits(s string) bool {
	size := len(s)
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = int(s[i] - '0')
	}
	for size > 2 {
		for i := 0; i < size-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		size--
	}
	return nums[0] == nums[1]
}

func HasSameDigits(s string) bool {
	return hasSameDigits(s)
}
