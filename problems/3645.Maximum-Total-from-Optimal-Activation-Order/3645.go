package leetcode

import "container/heap"

// 3645. Maximum Total from Optimal Activation Order
// You are given two integer arrays value and limit, both of length n.
// Initially, all elements are inactive. You may activate them in any order.
// To activate an inactive element at index i, the number of currently active elements must be strictly less than limit[i].
// When you activate the element at index i, it adds value[i] to the total activation value (i.e., the sum of value[i] for all elements that have undergone activation operations).
// After each activation, if the number of currently active elements becomes x, then all elements j with limit[j] <= x become permanently inactive, even if they are already active.
// Return the maximum total you can obtain by choosing the activation order optimally.

// MinHeap implements a min-heap of integers
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxTotal(value []int, limit []int) int64 {
	n := len(value)

	// Group items by their limit values
	limitGroups := make(map[int][]int)
	for i := 0; i < n; i++ {
		limitGroups[limit[i]] = append(limitGroups[limit[i]], value[i])
	}

	var totalSum int64

	// Process each group independently
	for limitVal, values := range limitGroups {
		// Use a min-heap of capacity limitVal to find top limitVal values
		h := &MinHeap{}
		heap.Init(h)

		for _, val := range values {
			if h.Len() < limitVal {
				// Heap has space, just add the value
				heap.Push(h, val)
			} else if h.Len() > 0 && val > (*h)[0] {
				// Heap is full, but current value is larger than the smallest in heap
				heap.Pop(h)
				heap.Push(h, val)
			}
		}

		// Sum all values in the heap (these are the top min(limitVal, len(values)) values)
		for h.Len() > 0 {
			totalSum += int64(heap.Pop(h).(int))
		}
	}

	return totalSum
}
