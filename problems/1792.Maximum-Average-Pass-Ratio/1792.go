package leetcode

import (
	"container/heap"
)

type ClassHeap [][]float64

func (h ClassHeap) Len() int {
	return len(h)
}
func (h ClassHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0] // max heap
}
func (h ClassHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ClassHeap) Push(x interface{}) {
	*h = append(*h, x.([]float64))
}

func (h *ClassHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func calcGain(pass, total float64) float64 {
	return (pass+1)/(total+1) - pass/total
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &ClassHeap{}
	heap.Init(h)

	// Initialize heap with gain, pass, total for each class
	for _, class := range classes {
		pass := float64(class[0])
		total := float64(class[1])
		gain := calcGain(pass, total)
		heap.Push(h, []float64{gain, pass, total})
	}

	// Distribute extra students
	for i := 0; i < extraStudents; i++ {
		top := heap.Pop(h).([]float64)
		pass := top[1] + 1
		total := top[2] + 1
		newGain := calcGain(pass, total)
		heap.Push(h, []float64{newGain, pass, total})
	}

	// Calculate average pass ratio
	totalRatio := 0.0
	for _, item := range *h {
		pass := item[1]
		total := item[2]
		totalRatio += pass / total
	}

	return totalRatio / float64(len(classes))
}
