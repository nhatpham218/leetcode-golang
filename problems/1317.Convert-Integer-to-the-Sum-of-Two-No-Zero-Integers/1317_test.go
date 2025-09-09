package leetcode

import (
	"testing"
)

func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool // We check if the result is valid rather than exact match
	}{
		{
			name:     "Example 1: n = 2",
			input:    2,
			expected: true,
		},
		{
			name:     "Example 2: n = 11",
			input:    11,
			expected: true,
		},
		{
			name:     "Small number: n = 3",
			input:    3,
			expected: true,
		},
		{
			name:     "Number with zeros: n = 10",
			input:    10,
			expected: true,
		},
		{
			name:     "Number with zeros: n = 100",
			input:    100,
			expected: true,
		},
		{
			name:     "Number with zeros: n = 1000",
			input:    1000,
			expected: true,
		},
		{
			name:     "Large number: n = 9999",
			input:    9999,
			expected: true,
		},
		{
			name:     "Maximum constraint: n = 10000",
			input:    10000,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getNoZeroIntegers(tt.input)

			// Check if result is not nil
			if result == nil {
				t.Errorf("getNoZeroIntegers(%d) = nil, want valid result", tt.input)
				return
			}

			// Check if result has exactly 2 elements
			if len(result) != 2 {
				t.Errorf("getNoZeroIntegers(%d) = %v, want array of length 2", tt.input, result)
				return
			}

			a, b := result[0], result[1]

			// Check if a + b = n
			if a+b != tt.input {
				t.Errorf("getNoZeroIntegers(%d) = [%d, %d], sum = %d, want sum = %d", tt.input, a, b, a+b, tt.input)
				return
			}

			// Check if both a and b are no-zero integers
			if !hasNoZero(a) {
				t.Errorf("getNoZeroIntegers(%d) = [%d, %d], %d contains zero", tt.input, a, b, a)
				return
			}

			if !hasNoZero(b) {
				t.Errorf("getNoZeroIntegers(%d) = [%d, %d], %d contains zero", tt.input, a, b, b)
				return
			}

			// Check if both a and b are positive
			if a <= 0 || b <= 0 {
				t.Errorf("getNoZeroIntegers(%d) = [%d, %d], both numbers should be positive", tt.input, a, b)
				return
			}
		})
	}
}

func TestHasNoZero(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Single digit no zero",
			input:    1,
			expected: true,
		},
		{
			name:     "Single digit with zero",
			input:    0,
			expected: false,
		},
		{
			name:     "Multiple digits no zero",
			input:    123,
			expected: true,
		},
		{
			name:     "Multiple digits with zero in middle",
			input:    105,
			expected: false,
		},
		{
			name:     "Multiple digits with zero at end",
			input:    120,
			expected: false,
		},
		{
			name:     "All nines",
			input:    999,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasNoZero(tt.input)
			if result != tt.expected {
				t.Errorf("hasNoZero(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
