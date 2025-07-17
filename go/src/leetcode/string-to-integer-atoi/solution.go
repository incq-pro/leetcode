// 8. 字符串转换整数 (atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/

package string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	i := 0
	n := len(s)
	// 去掉空格
	for ; i < n; i++ {
		if s[i] != ' ' {
			break
		}
	}
	// 判断正负号
	sig := 1
	if i < n {
		if s[i] == '-' {
			sig = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	// 去掉前导0
	for ; i < n; i++ {
		if s[i] != '0' {
			break
		}
	}
	// 开始计算值
	var r int64 = 0
	for ; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			r = r*10 + int64(s[i]-'0')
			if sig > 0 && r > math.MaxInt32 {
				return math.MaxInt32
			} else if sig < 0 && -r < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	return int(r) * sig
}

func MyAtoi(s string) int {
	return myAtoi(s)
}
