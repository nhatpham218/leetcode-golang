package leetcode

import (
	"math"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected float64
	}{
		{
			name:     "Example 1",
			points:   [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			expected: 2.00000,
		},
		{
			name:     "Example 2",
			points:   [][]int{{1, 0}, {0, 0}, {0, 1}},
			expected: 0.50000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestTriangleArea(tt.points)
			if math.Abs(result-tt.expected) > 1e-5 {
				t.Errorf("largestTriangleArea() = %v, want %v", result, tt.expected)
			}
		})
	}
}
