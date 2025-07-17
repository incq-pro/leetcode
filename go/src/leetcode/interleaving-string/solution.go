// 97. 交错字符串
// https://leetcode.cn/problems/interleaving-string/

package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if len(s3) != (m + n) {
		return false
	}
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 1; i <= m; i++ {
		if s1[i-1] == s3[i-1] {
			f[i][0] = true
		} else {
			break
		}
	}
	for j := 1; j <= n; j++ {
		if s2[j-1] == s3[j-1] {
			f[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if f[i-1][j] && s1[i-1] == s3[i+j-1] {
				f[i][j] = true
			} else if f[i][j-1] && s2[j-1] == s3[i+j-1] {
				f[i][j] = true
			}
		}
	}
	return f[m][n]
}

func IsInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleave(s1, s2, s3)
}
