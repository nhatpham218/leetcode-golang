package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    9669,
			expected: 9969,
		},
		{
			name:     "Example 2",
			input:    9996,
			expected: 9999,
		},
		{
			name:     "Example 3",
			input:    9999,
			expected: 9999,
		},
		{
			name:     "Single digit 6",
			input:    6,
			expected: 9,
		},
		{
			name:     "Single digit 9",
			input:    9,
			expected: 9,
		},
		{
			name:     "All 6s",
			input:    666,
			expected: 966,
		},
		{
			name:     "All 9s",
			input:    999,
			expected: 999,
		},
		{
			name:     "Mixed digits with 6 first",
			input:    6699,
			expected: 9699,
		},
		{
			name:     "Mixed digits with 9 first",
			input:    9966,
			expected: 9996,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum69Number(tt.input)
			assert.Equal(t, tt.expected, result, "maximum69Number(%d) should return %d, got %d", tt.input, tt.expected, result)
		})
	}
}
