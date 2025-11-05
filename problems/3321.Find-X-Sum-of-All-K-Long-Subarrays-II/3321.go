package leetcode

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

func findXSum(nums []int, k int, x int) []int64 {
	tracker := newXSumTracker(x)
	result := []int64{}

	for i := 0; i < len(nums); i++ {
		tracker.add(nums[i])
		if i >= k {
			tracker.remove(nums[i-k])
		}
		if i >= k-1 {
			result = append(result, tracker.getSum())
		}
	}

	return result
}

type xSumTracker struct {
	x         int
	sum       int64
	topX      *redblacktree.Tree[pair, struct{}]
	remaining *redblacktree.Tree[pair, struct{}]
	frequency map[int]int
}

type pair struct {
	freq int
	num  int
}

func comparePairs(a, b pair) int {
	if a.freq != b.freq {
		return a.freq - b.freq
	}
	return a.num - b.num
}

func newXSumTracker(x int) *xSumTracker {
	return &xSumTracker{
		x:         x,
		sum:       0,
		topX:      redblacktree.NewWith[pair, struct{}](comparePairs),
		remaining: redblacktree.NewWith[pair, struct{}](comparePairs),
		frequency: make(map[int]int),
	}
}

func (t *xSumTracker) add(num int) {
	if t.frequency[num] > 0 {
		t.removePair(pair{freq: t.frequency[num], num: num})
	}
	t.frequency[num]++
	t.addPair(pair{freq: t.frequency[num], num: num})
}

func (t *xSumTracker) remove(num int) {
	t.removePair(pair{freq: t.frequency[num], num: num})
	t.frequency[num]--
	if t.frequency[num] > 0 {
		t.addPair(pair{freq: t.frequency[num], num: num})
	}
}

func (t *xSumTracker) getSum() int64 {
	return t.sum
}

func (t *xSumTracker) addPair(p pair) {
	if t.topX.Size() < t.x {
		t.sum += int64(p.freq) * int64(p.num)
		t.topX.Put(p, struct{}{})
		return
	}

	minTopX := t.topX.Left().Key
	if comparePairs(p, minTopX) > 0 {
		t.sum += int64(p.freq) * int64(p.num)
		t.topX.Put(p, struct{}{})

		toMove := t.topX.Left().Key
		t.sum -= int64(toMove.freq) * int64(toMove.num)
		t.topX.Remove(toMove)
		t.remaining.Put(toMove, struct{}{})
	} else {
		t.remaining.Put(p, struct{}{})
	}
}

func (t *xSumTracker) removePair(p pair) {
	if _, found := t.topX.Get(p); found {
		t.sum -= int64(p.freq) * int64(p.num)
		t.topX.Remove(p)

		if t.remaining.Size() > 0 {
			maxRemaining := t.remaining.Right().Key
			t.sum += int64(maxRemaining.freq) * int64(maxRemaining.num)
			t.remaining.Remove(maxRemaining)
			t.topX.Put(maxRemaining, struct{}{})
		}
	} else if _, found := t.remaining.Get(p); found {
		t.remaining.Remove(p)
	}
}
