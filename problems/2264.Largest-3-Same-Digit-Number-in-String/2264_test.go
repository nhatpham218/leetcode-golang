package leetcode

import (
	"testing"
)

func TestLargestGoodInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "6777133339",
			expected: "777",
		},
		{
			name:     "Example 2",
			input:    "2300019",
			expected: "000",
		},
		{
			name:     "Example 3",
			input:    "42352338",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestGoodInteger(tt.input)
			if result != tt.expected {
				t.Errorf("largestGoodInteger(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
