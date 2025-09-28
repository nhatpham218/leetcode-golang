package leetcode

import "testing"

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 1, 2},
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 1, 10},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestPerimeter(tt.nums)
			if result != tt.expected {
				t.Errorf("largestPerimeter() = %v, want %v", result, tt.expected)
			}
		})
	}
}
