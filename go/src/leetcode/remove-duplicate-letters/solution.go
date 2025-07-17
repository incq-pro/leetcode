// 316. 去除重复字母
// https://leetcode.cn/problems/remove-duplicate-letters

package remove_duplicate_letters

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1]
				if left[last-'a'] <= 0 {
					break
				}
				stack = stack[0 : len(stack)-1]
				inStack[last-'a'] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}

func RemoveDuplicateLetters(s string) string {
	return removeDuplicateLetters(s)
}
