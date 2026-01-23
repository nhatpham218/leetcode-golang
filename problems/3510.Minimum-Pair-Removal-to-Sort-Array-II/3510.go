package leetcode

import "container/heap"

// 3510. Minimum Pair Removal to Sort Array II
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-ii/description/
func minimumPairRemoval(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Use int64 for values since sums can get very large
	values := make([]int64, n)
	for i, v := range nums {
		values[i] = int64(v)
	}

	// Doubly linked list to find prev/next active indices
	prev := make([]int, n)
	next := make([]int, n)
	removed := make([]bool, n)

	for i := 0; i < n; i++ {
		prev[i] = i - 1
		next[i] = i + 1
	}

	// Count "bad" pairs where values[i] > values[next[i]]
	badCount := 0
	for i := 0; i < n-1; i++ {
		if values[i] > values[i+1] {
			badCount++
		}
	}

	if badCount == 0 {
		return 0
	}

	// Min-heap of (sum, leftIndex) - minimum sum, then leftmost index
	h := &minHeap3510{}
	heap.Init(h)

	for i := 0; i < n-1; i++ {
		heap.Push(h, pair3510{sum: values[i] + values[i+1], left: i})
	}

	ops := 0

	for badCount > 0 {
		// Find valid minimum sum pair (lazy deletion for stale entries)
		var p pair3510
		for {
			p = heap.Pop(h).(pair3510)
			i := p.left
			// Pair is valid if: not removed, has next neighbor, and sum matches current values
			if !removed[i] && next[i] < n && p.sum == values[i]+values[next[i]] {
				break
			}
		}

		i := p.left
		j := next[i]
		pi := prev[i]
		nj := next[j]

		// Update badCount: remove old bad pairs involving (pi,i), (i,j), (j,nj)
		if pi >= 0 && values[pi] > values[i] {
			badCount--
		}
		if values[i] > values[j] {
			badCount--
		}
		if nj < n && values[j] > values[nj] {
			badCount--
		}

		// Merge: set values[i] = values[i] + values[j], mark j as removed
		values[i] = values[i] + values[j]
		removed[j] = true

		// Update linked list pointers
		next[i] = nj
		if nj < n {
			prev[nj] = i
		}

		// Update badCount: add new bad pairs involving (pi,i) and (i,nj)
		if pi >= 0 && values[pi] > values[i] {
			badCount++
		}
		if nj < n && values[i] > values[nj] {
			badCount++
		}

		// Add new pairs to heap
		if pi >= 0 {
			heap.Push(h, pair3510{sum: values[pi] + values[i], left: pi})
		}
		if nj < n {
			heap.Push(h, pair3510{sum: values[i] + values[nj], left: i})
		}

		ops++
	}

	return ops
}

type pair3510 struct {
	sum  int64
	left int
}

type minHeap3510 []pair3510

func (h minHeap3510) Len() int { return len(h) }
func (h minHeap3510) Less(i, j int) bool {
	if h[i].sum != h[j].sum {
		return h[i].sum < h[j].sum
	}
	return h[i].left < h[j].left // leftmost for ties
}
func (h minHeap3510) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap3510) Push(x any) {
	*h = append(*h, x.(pair3510))
}

func (h *minHeap3510) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
