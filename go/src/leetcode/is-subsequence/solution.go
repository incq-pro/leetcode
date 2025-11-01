// 392. 判断子序列
// https://leetcode.cn/problems/is-subsequence/

package is_subsequence

func isSubsequence(s string, t string) bool {
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func isSubsequence2(s string, t string) bool {
	n := len(t)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	for i, c := range t {
		for j := i; j >= 0; j-- {
			if dp[j][c-'a'] < 0 {
				dp[j][c-'a'] = i
			} else {
				break
			}
		}
	}

	k := 0
	for _, c := range s {
		next := dp[k][c-'a']
		if next < 0 {
			return false
		}
		k = next + 1
	}
	return true
}

func IsSubsequence(s string, t string) bool {
	return isSubsequence2(s, t)
}
