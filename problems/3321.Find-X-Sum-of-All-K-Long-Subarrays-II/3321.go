package leetcode

import "container/heap"

type pair struct {
	freq, val int
}

type maxHeap []pair

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val > h[j].val
}
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type minHeap []pair

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}
	return h[i].val < h[j].val
}
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findXSum(nums []int, k int, x int) []int64 {
	freq := make(map[int]int)
	top := &minHeap{}
	rest := &maxHeap{}
	heap.Init(top)
	heap.Init(rest)

	topMap := make(map[pair]int)
	restMap := make(map[pair]int)
	topSum := int64(0)
	topCnt := 0
	ans := make([]int64, 0, len(nums)-k+1)

	cleanTop := func() {
		for top.Len() > 0 && topMap[(*top)[0]] == 0 {
			heap.Pop(top)
		}
	}

	cleanRest := func() {
		for rest.Len() > 0 && restMap[(*rest)[0]] == 0 {
			heap.Pop(rest)
		}
	}

	balance := func() {
		for topCnt < x && len(restMap) > 0 {
			cleanRest()
			if rest.Len() == 0 {
				break
			}
			p := heap.Pop(rest).(pair)
			restMap[p]--
			topMap[p]++
			heap.Push(top, p)
			topSum += int64(p.freq) * int64(p.val)
			topCnt++
		}
		for topCnt > x {
			cleanTop()
			p := heap.Pop(top).(pair)
			topMap[p]--
			restMap[p]++
			heap.Push(rest, p)
			topSum -= int64(p.freq) * int64(p.val)
			topCnt--
		}
		for len(restMap) > 0 && topCnt > 0 {
			cleanRest()
			cleanTop()
			if rest.Len() == 0 || top.Len() == 0 {
				break
			}
			maxRest := (*rest)[0]
			minTop := (*top)[0]
			if maxRest.freq > minTop.freq || (maxRest.freq == minTop.freq && maxRest.val > minTop.val) {
				heap.Pop(rest)
				heap.Pop(top)
				restMap[maxRest]--
				topMap[minTop]--
				topMap[maxRest]++
				restMap[minTop]++
				heap.Push(top, maxRest)
				heap.Push(rest, minTop)
				topSum += int64(maxRest.freq)*int64(maxRest.val) - int64(minTop.freq)*int64(minTop.val)
			} else {
				break
			}
		}
	}

	add := func(num int) {
		old := pair{freq[num], num}
		if topMap[old] > 0 {
			topMap[old]--
			topSum -= int64(old.freq) * int64(old.val)
			topCnt--
		} else if restMap[old] > 0 {
			restMap[old]--
		}
		freq[num]++
		newP := pair{freq[num], num}
		restMap[newP]++
		heap.Push(rest, newP)
		balance()
	}

	remove := func(num int) {
		old := pair{freq[num], num}
		if topMap[old] > 0 {
			topMap[old]--
			topSum -= int64(old.freq) * int64(old.val)
			topCnt--
		} else {
			restMap[old]--
		}
		freq[num]--
		if freq[num] > 0 {
			newP := pair{freq[num], num}
			restMap[newP]++
			heap.Push(rest, newP)
		} else {
			delete(freq, num)
		}
		balance()
	}

	for i := 0; i < k; i++ {
		add(nums[i])
	}
	ans = append(ans, topSum)

	for i := k; i < len(nums); i++ {
		remove(nums[i-k])
		add(nums[i])
		ans = append(ans, topSum)
	}

	return ans
}
