// 3227. 字符串元音游戏
// https://leetcode.cn/problems/vowels-game-in-a-string

package vowels_game_in_a_string

func doesAliceWin(s string) bool {
	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for _, r := range s {
		if vowel[r] {
			return true
		}
	}
	return false
}

func DoesAliceWin(s string) bool {
	return doesAliceWin(s)
}
