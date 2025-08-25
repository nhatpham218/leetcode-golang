package leetcode

import (
	"testing"
)

func TestJudgePoint24(t *testing.T) {
	tests := []struct {
		name     string
		cards    []int
		expected bool
	}{
		{
			name:     "Example 1",
			cards:    []int{4, 1, 8, 7},
			expected: true,
		},
		{
			name:     "Example 2",
			cards:    []int{1, 2, 1, 2},
			expected: false,
		},
		{
			name:     "All same numbers - 1s",
			cards:    []int{1, 1, 1, 1},
			expected: false,
		},
		{
			name:     "All same numbers - 9s",
			cards:    []int{9, 9, 9, 9},
			expected: false,
		},
		{
			name:     "Simple case - 3,3,8,8",
			cards:    []int{3, 3, 8, 8},
			expected: true,
		},
		{
			name:     "Another valid case - 4,1,8,7",
			cards:    []int{4, 1, 8, 7},
			expected: true,
		},
		{
			name:     "Edge case - 1,2,3,4",
			cards:    []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "Edge case - 5,5,5,5",
			cards:    []int{5, 5, 5, 5},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := judgePoint24(tt.cards)
			if result != tt.expected {
				t.Errorf("judgePoint24(%v) = %v, want %v", tt.cards, result, tt.expected)
			}
		})
	}
}
