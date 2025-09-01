package leetcode

import (
	"math"
	"testing"
)

func TestMaxAverageRatio(t *testing.T) {
	tests := []struct {
		name          string
		classes       [][]int
		extraStudents int
		expected      float64
	}{
		{
			name:          "Example 1",
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			expected:      0.78333,
		},
		{
			name:          "Example 2",
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			expected:      0.53485,
		},
		{
			name:          "Single class",
			classes:       [][]int{{1, 2}},
			extraStudents: 1,
			expected:      0.66667,
		},
		{
			name:          "No extra students",
			classes:       [][]int{{1, 2}, {3, 5}},
			extraStudents: 0,
			expected:      0.55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxAverageRatio(tt.classes, tt.extraStudents)
			if math.Abs(result-tt.expected) > 1e-5 {
				t.Errorf("maxAverageRatio() = %v, want %v", result, tt.expected)
			}
		})
	}
}
