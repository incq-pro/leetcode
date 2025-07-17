//3136. 有效单词
// https://leetcode.cn/problems/valid-word

package valid_word

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	y, f := false, false
	for _, c := range word {
		if c >= '0' && c <= '9' {
		} else if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
				y = true
			} else {
				f = true
			}
		} else {
			return false
		}
	}
	return y && f
}

func IsValid(word string) bool {
	return isValid(word)
}
