// 3516. 找到最近的人
// https://leetcode.cn/problems/find-closest-person/

package find_closest_person

func findClosest(x int, y int, z int) int {
	delta := abs(x-z) - abs(y-z)
	if delta < 0 {
		return 1
	} else if delta == 0 {
		return 0
	} else {
		return 2
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func FindClosest(x int, y int, z int) int {
	return findClosest(x, y, z)
}
