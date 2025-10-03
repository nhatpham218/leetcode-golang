package leetcode

import (
	"container/heap"
)

// Cell represents a position in the matrix
type Cell struct {
	row, col int
	height   int
}

// MinHeap of cells for boundary processing
type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 407. Trapping Rain Water II
// https://leetcode.com/problems/trapping-rain-water-ii/
func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m <= 2 {
		return 0
	}
	n := len(heightMap[0])
	if n <= 2 {
		return 0
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &MinHeap{}
	heap.Init(h)

	// Add boundary cells to heap
	for i := 0; i < m; i++ {
		heap.Push(h, Cell{i, 0, heightMap[i][0]})
		heap.Push(h, Cell{i, n - 1, heightMap[i][n-1]})
		visited[i][0] = true
		visited[i][n-1] = true
	}

	for j := 0; j < n; j++ {
		heap.Push(h, Cell{0, j, heightMap[0][j]})
		heap.Push(h, Cell{m - 1, j, heightMap[m-1][j]})
		visited[0][j] = true
		visited[m-1][j] = true
	}

	maxHeight := 0
	totalWater := 0

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		maxHeight = max(maxHeight, cell.height)

		for _, dir := range directions {
			newRow := cell.row + dir[0]
			newCol := cell.col + dir[1]

			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				cellHeight := heightMap[newRow][newCol]

				if cellHeight < maxHeight {
					totalWater += maxHeight - cellHeight
				}

				heap.Push(h, Cell{newRow, newCol, cellHeight})
			}
		}
	}

	return totalWater
}
