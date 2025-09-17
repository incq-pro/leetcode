// 2197. 替换数组中的非互质数
// https://leetcode.cn/problems/replace-non-coprime-numbers-in-array

package replace_non_coprime_numbers_in_array

func replaceNonCoprimes(nums []int) []int {
	// l 表示长度
	l := len(nums)
	if l == 1 {
		return nums
	}
	// 最终结果
	result := make([]int, l)
	for {
		found := false
		j := 0
		for i := 0; i < l-1; {
			g := gcd(nums[i], nums[i+1])
			if g == 1 {
				result[j] = nums[i]
				j++
				i++
			} else {
				nums[i+1] = lcm1(nums[i], nums[i+1], g)
				i++
				found = true
				for k := j - 1; k >= 0; k-- {
					g2 := gcd(nums[i], result[k])
					if g2 != 1 {
						j--
						nums[i] = lcm1(nums[i], result[k], g2)
					} else {
						break
					}
				}
			}
		}
		result[j] = nums[l-1]
		j++
		l = j
		if !found {
			break
		}
		nums, result = result, nums
	}
	return result[:l]
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		b, a = a%b, b
	}
	return a
}

func lcm1(a, b, g int) int {
	return a * b / g
}

func ReplaceNonCoprimes(nums []int) []int {
	return replaceNonCoprimes(nums)
}
