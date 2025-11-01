// 2125. 银行中的激光束数量
// https://leetcode.cn/problems/number-of-laser-beams-in-a-bank

package number_of_laser_beams_in_a_bank

import "strings"

func numberOfBeams(bank []string) int {
	ans := 0
	last := 0
	for _, r := range bank {
		num := strings.Count(r, "1")
		if num > 0 {
			if last > 0 {
				ans += last * num
			}
			last = num
		}
	}
	return ans
}

func NumberOfBeams(bank []string) int {
	return numberOfBeams(bank)
}
