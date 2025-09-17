// 1935. 可以输入的最大单词数
// https://leetcode.cn/problems/maximum-number-of-words-you-can-type

package maximum_number_of_words_you_can_type

func canBeTypedWords(text string, brokenLetters string) int {
	b := make(map[rune]bool, len(brokenLetters))
	for _, c := range brokenLetters {
		b[c] = true
	}
	num := 0
	broken := false
	for _, c := range text {
		if b[c] {
			broken = true
		}
		if c == ' ' {
			if !broken {
				num++
			} else {
				broken = false
			}
		}
	}
	if !broken {
		num++
	}
	return num
}

func CanBeTypedWords(text string, brokenLetters string) int {
	return canBeTypedWords(text, brokenLetters)
}
