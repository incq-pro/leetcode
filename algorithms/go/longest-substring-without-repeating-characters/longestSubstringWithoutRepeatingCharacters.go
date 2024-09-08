// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package longest_substring_without_repeating_characters

import "fmt"

func lengthOfLongestSubstring(s string) int {
	index := make(map[rune]int, len(s))
	result := 0
	start := 0
	for i, ch := range s {
		preI, ok := index[ch]
		if ok && preI >= start {
			length := i - start
			start = preI + 1
			if result < length {
				result = length
			}
		}
		index[ch] = i
	}
	l := len(s) - start
	if result < l {
		result = l
	}
	return result
}

func LengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	return lengthOfLongestSubstring(s)
}
