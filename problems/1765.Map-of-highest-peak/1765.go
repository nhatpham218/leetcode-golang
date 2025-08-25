package leetcode

// highestPeak assigns heights to a map of land and water cells to maximize the highest peak.
//
// Problem Analysis:
// - Water cells (isWater[i][j] == 1) must have height 0
// - Land cells (isWater[i][j] == 0) can have any non-negative height
// - Adjacent cells can differ by at most 1 in height
// - Goal: Maximize the maximum height in the entire matrix
//
// Algorithm: Multi-source BFS (Breadth-First Search)
// The key insight is that to maximize the maximum height, we want each land cell
// to have the minimum possible height while satisfying constraints. This is achieved
// by using BFS starting from all water cells simultaneously.
//
// Time Complexity: O(m * n) where m and n are matrix dimensions
// Space Complexity: O(m * n) for the result matrix and queue
func highestPeak(isWater [][]int) [][]int {
	// Get matrix dimensions
	m, n := len(isWater), len(isWater[0])

	// Initialize result matrix with -1 (indicating unvisited cells)
	// We use -1 to distinguish between unvisited cells and cells with height 0
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = -1
		}
	}

	// Queue for BFS: stores [row, col] coordinates
	// We use a slice as queue for simplicity (could use container/list for better performance)
	queue := make([][]int, 0)

	// Step 1: Initialize all water cells with height 0 and add them to queue
	// This creates multiple starting points for our BFS (multi-source BFS)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				result[i][j] = 0 // Water cells must have height 0
				queue = append(queue, []int{i, j})
			}
		}
	}

	// Directions for 4-directional movement (up, down, left, right)
	// These represent the adjacent cells we need to check
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Step 2: BFS to assign heights to land cells
	// Process queue level by level, ensuring each level has height = previous level + 1
	for len(queue) > 0 {
		// Get current cell from front of queue
		current := queue[0]
		queue = queue[1:] // Remove processed element

		row, col := current[0], current[1]
		currentHeight := result[row][col]

		// Check all 4 adjacent cells
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			// Boundary check: ensure we don't go outside matrix bounds
			if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n {
				continue
			}

			// Skip if this cell has already been visited (assigned a height)
			if result[newRow][newCol] != -1 {
				continue
			}

			// Assign height to this land cell
			// Height = current cell's height + 1 (satisfies adjacent constraint)
			result[newRow][newCol] = currentHeight + 1

			// Add this cell to queue for processing its neighbors in next iteration
			queue = append(queue, []int{newRow, newCol})
		}
	}

	// Step 3: Handle edge case where there are no water cells
	// In this case, all cells are land and we should assign them the minimum height (1)
	// This satisfies the adjacent constraint and maximizes the maximum height given the constraint
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if result[i][j] == -1 {
				result[i][j] = 1 // Assign minimum positive height to unprocessed land cells
			}
		}
	}

	return result
}

// Why this algorithm works:
//
// 1. Correctness:
//    - Water cells get height 0 as required
//    - Each land cell gets the minimum possible height (distance from nearest water)
//    - Adjacent cells differ by exactly 1 (BFS level-by-level processing)
//
// 2. Optimality (maximizes maximum height):
//    - By giving each cell the minimum possible height, we maximize the height
//      of cells that are farthest from water
//    - The maximum height achieved is the maximum distance from any land cell to water
//
// 3. BFS guarantees:
//    - Cells are processed in order of increasing distance from water sources
//    - Each cell gets assigned the optimal (minimum) height for its position
