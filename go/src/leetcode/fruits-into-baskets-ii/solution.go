// 3477. 水果成篮 II
// https://leetcode.cn/problems/fruits-into-baskets-ii/

package fruits_into_baskets_ii

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	ans := 0
	for _, f := range fruits {
		found := false
		for i, b := range baskets {
			if b >= f {
				found = true
				baskets[i] = -1
				break
			}
		}
		if !found {
			ans++
		}
	}
	return ans
}

func NumOfUnplacedFruits(fruits []int, baskets []int) int {
	return numOfUnplacedFruits(fruits, baskets)
}
