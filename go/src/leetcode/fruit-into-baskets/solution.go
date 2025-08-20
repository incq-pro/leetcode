// 904. 水果成篮
// https://leetcode.cn/problems/fruit-into-baskets/

package fruit_into_baskets

func totalFruit(fruits []int) int {
	last := make(map[int]int)
	ans := 0
	s := 0
	e := 0
	set := map[int]bool{}
	for i, v := range fruits {
		last[v] = i
		e = i
		if ok, _ := set[v]; !ok {
			if len(set) < 2 {
				set[v] = true
			} else {
				for a := range set {
					if a != fruits[i-1] {
						ans = max(ans, e-s)
						delete(set, a)
						set[v] = true
						s = last[a] + 1
					}
				}
			}
		}
	}
	ans = max(ans, e-s+1)
	return ans
}

func TotalFruit(fruits []int) int {
	return totalFruit(fruits)
}
