// 3174. 清除数字
// https://leetcode.cn/problems/clear-digits

package clear_digits

import (
	"unicode"
	"unicode/utf8"
)

func clearDigits(s string) string {
	n := utf8.RuneCountInString(s)
	stack := make([]rune, n)
	top := 0
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			if top > 0 {
				top--
			}
		} else {
			stack[top] = ch
			top++
		}
	}
	return string(stack[0:top])
}

func ClearDigits(s string) string {
	return clearDigits(s)
}
