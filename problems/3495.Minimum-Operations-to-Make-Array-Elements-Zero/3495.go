package leetcode

// 3495. Minimum Operations to Make Array Elements Zero
// You are given a 2D array queries, where queries[i] is of the form [l, r]. Each queries[i] defines an array of integers nums consisting of elements ranging from l to r, both inclusive.
// In one operation, you can:
// Select two integers a and b from the array.
// Replace them with floor(a / 4) and floor(b / 4).
// Your task is to determine the minimum number of operations required to reduce all elements of the array to zero for each query. Return the sum of the results for all queries.

// Pre-computed cumulative operations at power-of-2 boundaries
var cumOps []int64
var boundaries []int64

func minOperations(queries [][]int) int64 {
	initRangeCache()
	var result int64

	for _, query := range queries {
		l, r := query[0], query[1]

		// Calculate total operations needed for range [l, r]
		totalOps := countOpsInRange(r) - countOpsInRange(l-1)

		// Convert to pair operations (ceiling division by 2)
		pairOps := (totalOps + 1) / 2

		result += pairOps
	}

	return result
}

// initRangeCache pre-computes cumulative operations at range boundaries
func initRangeCache() {
	if len(cumOps) > 0 {
		return // Already initialized
	}

	// Pre-compute for ranges up to 2^31 (covers constraint: r <= 10^9)
	cumOps = []int64{0} // cumOps[0] = operations for [1,0] = 0
	boundaries = []int64{0}

	var cumTotal int64
	for bitLen := 1; bitLen <= 31; bitLen++ {
		// Numbers with bitLen bits: [2^(bitLen-1), 2^bitLen - 1]
		start := int64(1) << (bitLen - 1) // 2^(bitLen-1)
		end := (int64(1) << bitLen) - 1   // 2^bitLen - 1

		// Count of numbers with this binary length
		count := end - start + 1

		// Operations needed for each number with bitLen bits
		opsPerNum := int64((bitLen + 1) / 2) // ceiling(bitLen / 2)

		cumTotal += count * opsPerNum

		boundaries = append(boundaries, end)
		cumOps = append(cumOps, cumTotal)
	}
}

// countOpsInRange calculates total operations needed for all numbers in [1, n]
func countOpsInRange(n int) int64 {
	if n <= 0 {
		return 0
	}

	// Find the range that n belongs to
	for i := 1; i < len(boundaries); i++ {
		if int64(n) <= boundaries[i] {
			// n is in range [boundaries[i-1]+1, boundaries[i]]
			// Get cumulative up to previous boundary
			result := cumOps[i-1]

			// Add operations for numbers in partial range
			rangeStart := boundaries[i-1] + 1
			if rangeStart <= int64(n) {
				count := int64(n) - rangeStart + 1
				opsPerNum := int64((i + 1) / 2) // ceiling(i / 2), i is bitLen
				result += count * opsPerNum
			}

			return result
		}
	}

	// If n is larger than our pre-computed ranges, should not happen
	return 0
}
