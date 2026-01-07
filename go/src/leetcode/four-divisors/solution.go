// 1390. 四因数
// https://leetcode.cn/problems/four-divisors

package four_divisors

import "fmt"

func sumFourDivisors(nums []int) int {
	// 构造质数表
	primes, nonePrimes := makePrimes(nums)
	ans := 0
	for _, num := range nums {
		for _, p := range primes {
			if p >= num {
				break
			}
			q := num / p
			if q*p == num {
				if p != q && (!nonePrimes[q] || q == p*p) {
					fmt.Println(num)
					ans += 1 + num + p + q
				}
				break
			}
		}
	}
	return ans
}

func makePrimes(nums []int) ([]int, []bool) {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 构造质数表
	all := make([]bool, m+1)
	for i := 2; i <= m; i++ {
		if all[i] {
			continue
		}
		for j := i * 2; j <= m; j += i {
			all[j] = true
		}
	}
	primes := make([]int, 0)
	for i := 2; i <= m; i++ {
		if !all[i] {
			primes = append(primes, i)
		}
	}
	return primes, all
}

func SumFourDivisors(nums []int) int {
	return sumFourDivisors(nums)
}
