// 2414. 最长的字母序连续子字符串的长度
// https://leetcode.cn/problems/length-of-the-longest-alphabetical-continuous-substring

package length_of_the_longest_alphabetical_continuous_substring

func longestContinuousSubstring(s string) int {
	base := 0
	r := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			r = max(r, i-base+1)
			if r >= 26 {
				return r
			}
		} else {
			base = i
		}
	}
	return r
}

func LongestContinuousSubstring(s string) int {
	return longestContinuousSubstring(s)
}
