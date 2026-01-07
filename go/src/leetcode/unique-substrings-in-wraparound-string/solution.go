// 467. 环绕字符串中唯一的子字符串
// https://leetcode.cn/problems/unique-substrings-in-wraparound-string

package unique_substrings_in_wraparound_string

func findSubstringInWraproundString(s string) int {
	n := len(s)
	dp := make([]int, 26)

	left, right := 0, 1
	for ; right < n; right++ {
		dp[s[right-1]-'a'] = max(dp[s[right-1]-'a'], right-left)
		if s[right]%26 != (s[right-1]+1)%26 {
			left = right
		}
	}
	if left < right {
		dp[s[right-1]-'a'] = max(dp[s[right-1]-'a'], right-left)
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}
	return ans
}

func FindSubstringInWraproundString(s string) int {
	return findSubstringInWraproundString(s)
}
