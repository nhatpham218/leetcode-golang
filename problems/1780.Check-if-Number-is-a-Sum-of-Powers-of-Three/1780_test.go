package main

import (
	"testing"
)

func Test_checkPowersOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    12,
			want: true,
		},
		{
			name: "Example 2",
			n:    91,
			want: true,
		},
		{
			name: "Example 3",
			n:    21,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.n); got != tt.want {
				t.Errorf("checkPowersOfThree(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
