// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/

package longest_palindromic_substring

func longestPalindrome(s string) string {
	span := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		pre := i - 1
		if pre-span[pre]-1 >= 0 && s[i] == s[pre-span[pre]-1] {
			span[i] = span[pre] + 2
			continue
		}
		for j := pre - span[pre]; j <= i; j++ {
			if isPalindrome(s, j, i) {
				span[i] = i - j
				break
			}
		}
	}
	maxI := 0
	for i, length := range span {
		if length > span[maxI] {
			maxI = i
		}
	}
	return s[maxI-span[maxI] : maxI+1]
}

func isPalindrome(s string, l int, h int) bool {
	for l < h {
		if s[l] != s[h] {
			return false
		}
		l++
		h--
	}
	return true
}

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}
