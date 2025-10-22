// 3408. 设计任务管理器
// https://leetcode.cn/problems/design-task-manager

package design_task_manager

import "container/heap"

type task struct {
	userId   int
	taskId   int
	priority int
}

type hp []*task

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId >= h[j].taskId
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(*task))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	m  map[int]*task
	pq hp
}

func Constructor(tasks [][]int) TaskManager {
	m := make(map[int]*task, len(tasks))
	pq := hp(make([]*task, len(tasks)))
	for i, t := range tasks {
		tt := task{t[0], t[1], t[2]}
		m[tt.taskId] = &tt
		pq[i] = &tt
	}
	heap.Init(&pq)
	return TaskManager{
		m:  m,
		pq: pq,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := task{userId, taskId, priority}
	this.m[t.taskId] = &t
	heap.Push(&this.pq, &t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	if t, ok := this.m[taskId]; ok && t.priority != newPriority {
		t = &task{t.userId, taskId, newPriority}
		this.m[taskId] = t
		heap.Push(&this.pq, t)
	}
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.m, taskId)
}

func (this *TaskManager) ExecTop() int {
	for this.pq.Len() > 0 {
		t := heap.Pop(&this.pq).(*task)
		if t2, ok := this.m[t.taskId]; !ok || t2 != t {
			continue
		}
		delete(this.m, t.taskId)
		return t.userId
	}
	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
