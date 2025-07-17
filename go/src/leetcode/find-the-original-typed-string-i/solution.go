// 3330. 找到初始输入字符串 I
// https://leetcode.cn/problems/find-the-original-typed-string-i/

package find_the_original_typed_string_i

func possibleStringCount(word string) int {
	count := 1
	ch := rune(word[0])
	i := 0
	for _, c := range word {
		if c == ch {
			i++
		} else {
			count += (i - 1)
			ch = c
			i = 1
		}
	}
	count += (i - 1)
	return count
}

func PossibleStringCount(word string) int {
	return possibleStringCount(word)
}
