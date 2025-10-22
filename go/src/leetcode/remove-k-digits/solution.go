// 402. 移掉 K 位数字
// https://leetcode.cn/problems/remove-k-digits/

package remove_k_digits

func removeKdigits(num string, k int) string {
	ans := []byte{num[0]}
	i := 1
	for j := 0; j < k && len(ans) > 0; j++ {
		for ; i < len(num) && ans[len(ans)-1] <= num[i]; i++ {
			ans = append(ans, num[i])
		}
		// 删除ans最后元素
		ans = ans[:len(ans)-1]
		if len(ans) == 0 && i < len(num) {
			for i < len(num) {
				if num[i] != '0' {
					ans = append(ans, num[i])
					i++
					break
				}
				i++
			}
		}
	}
	if len(ans) == 0 {
		return "0"
	}
	return string(ans) + num[i:]
}

func RemoveKdigits(num string, k int) string {
	return removeKdigits(num, k)
}
