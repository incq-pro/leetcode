// 3244. 新增道路查询后的最短距离 II
// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/

package shortest_distance_after_road_addition_queries_ii

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	res := make([]int, len(queries))
	distance := n - 1
	next := make([]int, n-1)
	for i := 0; i < len(next); i++ {
		next[i] = i + 1
	}
	for i, query := range queries {
		u, v := query[0], query[1]
		if next[u] > 0 && next[u] < v {
			reduce := 0
			j := next[u]
			next[u] = v
			for j > 0 && j < v {
				j, next[j] = next[j], -1
				reduce++
			}
			distance -= reduce
		}
		res[i] = distance
	}
	return res
}

func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	return shortestDistanceAfterQueries(n, queries)
}
