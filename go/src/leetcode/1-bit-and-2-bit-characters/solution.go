// 717. 1 比特与 2 比特字符
// https://leetcode.cn/problems/1-bit-and-2-bit-characters/

package __bit_and_2_bit_characters

func isOneBitCharacter(bits []int) bool {
	isOne := false
	n := len(bits)
	for i := 0; i < n; i++ {
		if bits[i] == 0 {
			isOne = true
		} else {
			isOne = false
			i++
		}
	}
	return isOne
}

func IsOneBitCharacter(bits []int) bool {
	return isOneBitCharacter(bits)
}
