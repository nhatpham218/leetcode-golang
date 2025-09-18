package leetcode

import (
	"container/heap"
)

// 3408. Design Task Manager
// https://leetcode.com/problems/design-task-manager/

type Task struct {
	userId   int
	taskId   int
	priority int
}

// MaxHeap for tasks (higher priority first, higher taskId for ties)
type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].taskId > h[j].taskId // higher taskId first for ties
	}
	return h[i].priority > h[j].priority // higher priority first
}
func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}
func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TaskManager struct {
	taskMap map[int]Task // taskId -> Task info
	heap    *TaskHeap    // priority queue for tasks
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		taskMap: make(map[int]Task),
		heap:    &TaskHeap{},
	}

	// Initialize with given tasks
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		tm.Add(userId, taskId, priority)
	}

	return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	task := Task{userId: userId, taskId: taskId, priority: priority}
	tm.taskMap[taskId] = task
	heap.Push(tm.heap, task)
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	// Update task in map
	task := tm.taskMap[taskId]
	task.priority = newPriority
	tm.taskMap[taskId] = task

	// Add updated task to heap (lazy deletion - old entries cleaned up in ExecTop)
	heap.Push(tm.heap, task)
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.taskMap, taskId)
}

func (tm *TaskManager) ExecTop() int {
	// Clean up stale entries from heap top
	for tm.heap.Len() > 0 {
		top := (*tm.heap)[0]

		// Check if task still exists and has current priority
		if currentTask, exists := tm.taskMap[top.taskId]; exists &&
			currentTask.priority == top.priority {
			// Execute this task
			userId := currentTask.userId
			delete(tm.taskMap, top.taskId)
			heap.Pop(tm.heap)
			return userId
		}

		// Remove stale entry
		heap.Pop(tm.heap)
	}

	return -1 // No tasks available
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
