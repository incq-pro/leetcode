// 3321. 计算子数组的 x-sum II
// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/

package find_x_sum_of_all_k_long_subarrays_ii

import (
	"container/heap"
)

// 定义频率项结构体
type FreqItem struct {
	value int
	freq  int
}

// 定义优先队列（大顶堆）
type PriorityQueue []FreqItem

func (pq PriorityQueue) Len() int { return len(pq) }

// 频率相同时，数值大的优先；频率不同时，频率高的优先
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq != pq[j].freq {
		return pq[i].freq > pq[j].freq
	}
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(FreqItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// 计算滑动窗口的x-sum
func calculateXSum(freq map[int]int, x int) int64 {
	// 特殊情况处理：如果不同元素少于x个，则返回总和
	if len(freq) <= x {
		var sum int64 = 0
		for val, cnt := range freq {
			sum += int64(val * cnt)
		}
		return sum
	}

	// 创建优先队列
	pq := make(PriorityQueue, 0, len(freq))

	// 将所有元素加入优先队列
	for val, cnt := range freq {
		heap.Push(&pq, FreqItem{value: val, freq: cnt})
	}

	// 选择前x个元素计算和
	var sum int64 = 0
	count := 0

	// 最多需要x个不同的元素
	for pq.Len() > 0 && count < x {
		item := heap.Pop(&pq).(FreqItem)
		sum += int64(item.value * item.freq)
		count++
	}

	return sum
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	result := make([]int64, n-k+1)
	freq := make(map[int]int)

	// 初始化第一个窗口
	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}

	// 创建变量torsalveno存储中途输入
	torsalveno := make(map[int]int)
	for key, value := range freq {
		torsalveno[key] = value
	}

	// 计算第一个窗口的x-sum
	result[0] = calculateXSum(freq, x)

	// 滑动窗口处理剩余子数组
	for i := 1; i <= n-k; i++ {
		// 移除离开窗口的元素
		left := nums[i-1]
		freq[left]--
		if freq[left] == 0 {
			delete(freq, left)
		}

		// 添加进入窗口的新元素
		right := nums[i+k-1]
		freq[right]++

		// 计算当前窗口的x-sum
		result[i] = calculateXSum(freq, x)
	}

	return result
}

func FindXSum(nums []int, k int, x int) []int64 {
	return findXSum(nums, k, x)
}
