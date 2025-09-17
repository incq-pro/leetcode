// 2785. 将字符串中的元音字母排序
// https://leetcode.cn/problems/sort-vowels-in-a-string/

package sort_vowels_in_a_string

import "sort"

func sortVowels(s string) string {
	b := []byte(s)
	index := make([]int, 0, len(b))
	vowels := make([]byte, 0, len(b))
	for i, c := range b {
		if isVowel(c) {
			index = append(index, i)
			vowels = append(vowels, c)
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	for i, j := range index {
		b[j] = vowels[i]
	}
	return string(b)
}

func isVowel(c byte) bool {
	return 'a' == c || 'e' == c || 'i' == c || 'o' == c || 'u' == c || 'A' == c || 'E' == c || 'I' == c || 'O' == c || 'U' == c
}

func SortVowels(s string) string {
	return sortVowels(s)
}
