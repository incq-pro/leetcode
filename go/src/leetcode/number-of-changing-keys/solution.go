// 3019. 按键变更的次数
// https://leetcode.cn/problems/number-of-changing-keys/

package number_of_changing_keys

import "strings"

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			count++
		}
	}
	return count
}
func CountKeyChanges(s string) int {
	return countKeyChanges(s)
}
