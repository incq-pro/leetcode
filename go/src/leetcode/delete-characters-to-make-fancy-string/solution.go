// 1957. 删除字符使字符串变好
// https://leetcode.cn/problems/delete-characters-to-make-fancy-string

package delete_characters_to_make_fancy_string

import "strings"

func makeFancyString(s string) string {
	c := 0
	var ch rune
	var builder strings.Builder
	for _, x := range s {
		if ch == x {
			c++
		} else {
			ch = x
			c = 1
		}
		if c < 3 {
			builder.WriteRune(x)

		}
	}
	return builder.String()
}

func MakeFancyString(s string) string {
	return makeFancyString(s)
}
