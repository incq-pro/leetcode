// 1317. 将整数转换为两个无零整数的和
// https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

package convert_integer_to_the_sum_of_two_no_zero_integers

func getNoZeroIntegers(n int) []int {
	mid := n / 2
	for i := 1; i <= mid; i++ {
		if noneZero(i) && noneZero(n-i) {
			return []int{i, n - i}
		}
	}
	return nil
}

func noneZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func GetNoZeroIntegers(n int) []int {
	return getNoZeroIntegers(n)
}
