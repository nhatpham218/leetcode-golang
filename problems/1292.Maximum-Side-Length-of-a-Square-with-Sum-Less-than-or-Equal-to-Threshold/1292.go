package leetcode

// 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/description/
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])

	// Precompute row prefix sums
	// rowPrefix[i][j] = sum of mat[i][0..j-1]
	rowPrefix := make([][]int, m)
	for i := 0; i < m; i++ {
		rowPrefix[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			rowPrefix[i][j+1] = rowPrefix[i][j] + mat[i][j]
		}
	}

	// Helper to get sum of row i from column c1 to c2 (inclusive)
	rowSum := func(i, c1, c2 int) int {
		return rowPrefix[i][c2+1] - rowPrefix[i][c1]
	}

	result := 0

	// Brute force: try all possible top-left corners and side lengths
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxK := min(m-i, n-j)
			// Only try squares larger than current result
			if maxK <= result {
				continue
			}

			// Incrementally compute square sum
			squareSum := 0
			for k := 1; k <= maxK; k++ {
				// Add the new bottom row (row i+k-1) from column j to j+k-1
				squareSum += rowSum(i+k-1, j, j+k-1)
				// Add the new right column cells (except bottom-right which is already added)
				// These are cells (i, j+k-1), (i+1, j+k-1), ..., (i+k-2, j+k-1)
				for row := i; row < i+k-1; row++ {
					squareSum += rowSum(row, j+k-1, j+k-1)
				}

				if squareSum <= threshold {
					result = max(result, k)
				} else if k > result {
					// If current size already exceeds threshold and is bigger than result,
					// larger sizes will also exceed, so break
					break
				}
			}
		}
	}

	return result
}
