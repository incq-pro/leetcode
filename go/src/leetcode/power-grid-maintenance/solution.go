// 3607. 电网维护
// https://leetcode.cn/problems/power-grid-maintenance/

package power_grid_maintenance

import "container/heap"

type H []int

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i] <= h[j]
}
func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *H) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	graph := make(map[int][]int, len(connections))
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	hps := make([]**H, c+1)
	for u, _ := range graph {
		if hps[u] != nil {
			continue
		}
		h := (H)([]int{})
		hp := &h
		queue := []int{u}
		for len(queue) > 0 {
			ele := queue[0]
			queue = queue[1:]
			hps[ele] = &hp
			heap.Push(&h, ele)
			hp = &h
			for _, vv := range graph[ele] {
				if hps[vv] == nil {
					queue = append(queue, vv)
				}
			}
		}
	}
	// 0 is online
	online := make([]int, c+1)
	var ans []int
	for _, q := range queries {
		op, num := q[0], q[1]
		if op == 2 {
			online[num] = 1
		} else {
			// 0 is online
			a := -1
			if online[num] == 0 {
				a = num
			} else {
				hp := hps[num]
				if hp != nil && *hp != nil {
					h := *hp
					for h.Len() > 0 {
						if online[(*h)[0]] == 0 {
							a = (*h)[0]
							break
						} else {
							heap.Pop(h)
						}
					}
				}
			}
			ans = append(ans, a)
		}
	}
	return ans
}

func ProcessQueries(c int, connections [][]int, queries [][]int) []int {
	return processQueries(c, connections, queries)
}

// 以下为官方题解

func processQueriesLc(c int, connections [][]int, queries [][]int) []int {
	graph := make([][]int, c+1)
	vertices := make([]Vertex, c+1)

	for i := 0; i <= c; i++ {
		graph[i] = make([]int, 0)
		vertices[i] = Vertex{vertexId: i, powerGridId: -1}
	}

	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	powerGrids := make([]*IntHeap, 0)
	for i, powerGridId := 1, 0; i <= c; i++ {
		v := &vertices[i]
		if v.powerGridId == -1 {
			powerGrid := &IntHeap{}
			heap.Init(powerGrid)
			traverse(v, powerGridId, powerGrid, graph, vertices)
			powerGrids = append(powerGrids, powerGrid)
			powerGridId++
		}
	}

	ans := make([]int, 0)
	for _, q := range queries {
		op, x := q[0], q[1]
		if op == 1 {
			if !vertices[x].offline {
				ans = append(ans, x)
			} else {
				powerGrid := powerGrids[vertices[x].powerGridId]
				for powerGrid.Len() > 0 && vertices[(*powerGrid)[0]].offline {
					heap.Pop(powerGrid)
				}
				if powerGrid.Len() > 0 {
					ans = append(ans, (*powerGrid)[0])
				} else {
					ans = append(ans, -1)
				}
			}
		} else if op == 2 {
			vertices[x].offline = true
		}
	}

	return ans
}

func traverse(u *Vertex, powerGridId int, powerGrid *IntHeap, graph [][]int, vertices []Vertex) {
	u.powerGridId = powerGridId
	heap.Push(powerGrid, u.vertexId)
	for _, vid := range graph[u.vertexId] {
		v := &vertices[vid]
		if v.powerGridId == -1 {
			traverse(v, powerGridId, powerGrid, graph, vertices)
		}
	}
}

type Vertex struct {
	vertexId    int
	offline     bool
	powerGridId int
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
