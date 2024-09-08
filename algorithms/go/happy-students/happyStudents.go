// 2860. 让所有学生保持开心的分组方法数
// https://leetcode.cn/problems/happy-students/

package happy_students

import "slices"

func countWays(nums []int) int {
	slices.Sort(nums)
	ways := 0
	// 待校验的index, 加入集合并通过，才算校验完成
	toCheck := 0
	// 加入集合的同学个数，默认是同学排序后依次加入
	for size := 0; size <= len(nums); size++ {
		// 校验已加入集合的同学是否满意
		for ; toCheck < size; toCheck++ {
			// 已加入的同学不满意
			if size <= nums[toCheck] {
				break
			}
		}
		// 已加入同学都满意，校验未加入的同学是否满意
		if toCheck == size {
			// 所有同学都已加入，满足条件
			if toCheck == len(nums) {
				ways++
			} else if size < nums[toCheck] {
				// 未加入的同学也满意，满足条件
				ways++
			}
		}
	}
	return ways
}

func CountWays(nums []int) int {
	return countWays(nums)
}
