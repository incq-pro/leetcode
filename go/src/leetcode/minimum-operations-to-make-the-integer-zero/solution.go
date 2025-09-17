// 2749. 得到整数零需要执行的最少操作数
// https://leetcode.cn/problems/minimum-operations-to-make-the-integer-zero/

package minimum_operations_to_make_the_integer_zero

func makeTheIntegerZero(num1 int, num2 int) int {
	for k := 1; ; k++ {
		x := int64(num1) - int64(num2)*int64(k)
		if int64(k) > x {
			return -1
		} else if k >= bitS(x) {
			return k
		}
	}
}

func bitS(n int64) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func MakeTheIntegerZero(num1 int, num2 int) int {
	return makeTheIntegerZero(num1, num2)
}
