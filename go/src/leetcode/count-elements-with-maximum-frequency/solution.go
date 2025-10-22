// 3005. 最大频率元素计数
// https://leetcode.cn/problems/count-elements-with-maximum-frequency/

package count_elements_with_maximum_frequency

func maxFrequencyElements(nums []int) int {
	feq := map[int]int{}
	mf := 0
	for _, num := range nums {
		f := feq[num] + 1
		feq[num] = f
		mf = max(mf, f)
	}
	s := 0
	for _, v := range feq {
		if v == mf {
			s += v
		}
	}
	return s
}

func MaxFrequencyElements(nums []int) int {
	return maxFrequencyElements(nums)

}
