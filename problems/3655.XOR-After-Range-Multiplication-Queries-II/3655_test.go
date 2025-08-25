package leetcode

import "testing"

func TestXorAfterQueries(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1},
			queries:  [][]int{{0, 2, 1, 4}},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 1, 5, 4},
			queries:  [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			expected: 31,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			queries:  [][]int{{0, 0, 1, 3}},
			expected: 15,
		},
		{
			name:     "No queries",
			nums:     []int{1, 2, 3},
			queries:  [][]int{},
			expected: 0, // 1 ^ 2 ^ 3 = 0
		},
		{
			name:     "Step size greater than range",
			nums:     []int{1, 2, 3, 4, 5},
			queries:  [][]int{{0, 4, 10, 2}},
			expected: 2, // Only index 0 is affected: [2, 2, 3, 4, 5] -> 2^2^3^4^5 = 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums to avoid modifying the original
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			result := xorAfterQueries(numsCopy, tt.queries)
			if result != tt.expected {
				t.Errorf("xorAfterQueries() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkXorAfterQueries(t *testing.B) {
	nums := []int{2, 3, 1, 5, 4}
	queries := [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}}

	for i := 0; i < t.N; i++ {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		xorAfterQueries(numsCopy, queries)
	}
}
