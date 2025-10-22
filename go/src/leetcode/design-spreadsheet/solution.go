// 3484. 设计电子表格
// https://leetcode.cn/problems/design-spreadsheet

package design_spreadsheet

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	d    map[string]int
	rows int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		d:    make(map[string]int),
		rows: rows,
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.d[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.d, cell)
}

func (this *Spreadsheet) GetValue(formula string) int {
	parts := strings.Split(formula[1:], "+")
	s := 0
	for _, p := range parts {
		if p[0] >= '0' && p[0] <= '9' {
			v, _ := strconv.Atoi(p)
			s += v
		} else {
			s += this.d[p]
		}
	}
	return s
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
