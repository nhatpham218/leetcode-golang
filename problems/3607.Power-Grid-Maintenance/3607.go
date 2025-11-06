package leetcode

import (
	"container/heap"
	"sort"
)

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

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return
	}
	if d.rank[px] < d.rank[py] {
		d.parent[px] = py
	} else if d.rank[px] > d.rank[py] {
		d.parent[py] = px
	} else {
		d.parent[py] = px
		d.rank[px]++
	}
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c + 1)

	for _, conn := range connections {
		dsu.Union(conn[0], conn[1])
	}

	compStations := make(map[int][]int)
	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		compStations[root] = append(compStations[root], i)
	}

	compHeaps := make(map[int]*MinHeap)
	for root, stations := range compStations {
		sort.Ints(stations)
		h := MinHeap(stations)
		heap.Init(&h)
		compHeaps[root] = &h
	}

	online := make([]bool, c+1)
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	result := []int{}

	for _, query := range queries {
		qType, station := query[0], query[1]

		if qType == 1 {
			if online[station] {
				result = append(result, station)
			} else {
				root := dsu.Find(station)
				h := compHeaps[root]

				for h.Len() > 0 && !online[(*h)[0]] {
					heap.Pop(h)
				}

				if h.Len() > 0 {
					result = append(result, (*h)[0])
				} else {
					result = append(result, -1)
				}
			}
		} else {
			online[station] = false
		}
	}

	return result
}
