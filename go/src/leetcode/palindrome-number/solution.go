// 9. 回文数
// https://leetcode.cn/problems/palindrome-number/

package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a int64 = int64(x)
	var y int64 = 0
	for a > 0 {
		m := a % 10
		a = a / 10
		y = y*10 + m
	}
	return y == int64(x)
}

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}
