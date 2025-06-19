// https://leetcode.cn/problems/count-prefixes-of-a-given-string
// 2255. 统计是给定字符串前缀的字符串数目

package count_prefixes_of_a_given_string

import "strings"

func countPrefixes(words []string, s string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			count++
		}
	}
	return count
}

func CountPrefixes(words []string, s string) int {
	return countPrefixes(words, s)
}
