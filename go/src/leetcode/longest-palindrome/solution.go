// 409. 最长回文串
// https://leetcode.cn/problems/longest-palindrome

package longest_palindrome

func longestPalindrome(s string) int {
	cnt := make(map[rune]int, len(s))
	for _, c := range s {
		cnt[c] += 1
	}
	ans := 0
	a := 0
	for _, v := range cnt {
		ans += v
		if v%2 != 0 {
			ans -= a
			a = 1
		}
	}
	return ans
}

func LongestPalindrome(s string) int {
	return longestPalindrome(s)
}
