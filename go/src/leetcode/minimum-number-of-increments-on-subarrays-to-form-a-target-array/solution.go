// 1526. 形成目标数组的子数组最少增加次数
// https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/

package minimum_number_of_increments_on_subarrays_to_form_a_target_array

func minNumberOperations(target []int) int {
	n := len(target)
	ans := target[0]
	for i := 1; i < n; i++ {
		ans += max(target[i]-target[i-1], 0)
	}
	return ans
}

// minNumberOperationsN2 O(n^2)
func minNumberOperationsN2(target []int) int {
	// 队列
	queue := [][]int{{0, len(target) - 1, 0}}
	ans := 0
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		left, right, base := task[0], task[1], task[2]
		minValueIndexes := []int{left}
		minValue := target[left]
		for i := left + 1; i <= right; i++ {
			if target[i] < minValue {
				minValue = target[i]
				minValueIndexes = []int{i}
			} else if target[i] == minValue {
				minValueIndexes = append(minValueIndexes, i)
			}
		}
		ans += minValue - base
		base = minValue
		for j, index := range minValueIndexes {
			if j == 0 {
				// first one
				if index > left {
					queue = append(queue, []int{left, index - 1, base})
				}
			}
			if j == len(minValueIndexes)-1 {
				// last one
				if index < right {
					queue = append(queue, []int{index + 1, right, base})
				}
			} else if minValueIndexes[j+1]-index > 1 {
				queue = append(queue, []int{index + 1, minValueIndexes[j+1] - 1, base})
			}
		}
	}
	return ans
}

func MinNumberOperations(target []int) int {
	return minNumberOperations(target)
}
