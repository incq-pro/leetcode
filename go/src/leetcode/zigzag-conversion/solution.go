// 6. Z 字形变换
// https://leetcode.cn/problems/zigzag-conversion

package zigzag_conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	m := 2*numRows - 2
	builders := make([]strings.Builder, numRows)
	for i, ch := range s {
		row := i % m
		if row >= numRows {
			row = m - row
		}
		builders[row].WriteRune(ch)
	}
	result := strings.Builder{}
	result.Grow(len(s))
	for _, b := range builders {
		result.WriteString(b.String())
	}
	return result.String()
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}
