package leetcode

import "testing"

func TestMinScoreTriangulation(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Example 1",
			values:   []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Example 2",
			values:   []int{3, 7, 4, 5},
			expected: 144,
		},
		{
			name:     "Example 3",
			values:   []int{1, 3, 1, 4, 1, 5},
			expected: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minScoreTriangulation(tt.values)
			if result != tt.expected {
				t.Errorf("minScoreTriangulation() = %v, want %v", result, tt.expected)
			}
		})
	}
}
