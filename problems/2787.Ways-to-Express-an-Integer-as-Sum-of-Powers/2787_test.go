package leetcode

import (
	"testing"
)

func Test_numberOfWays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		x    int
		want int
	}{
		{
			name: "Example 1",
			n:    10,
			x:    2,
			want: 1,
		},
		{
			name: "Example 2",
			n:    4,
			x:    1,
			want: 2,
		},
		{
			name: "tc3",
			n:    1,
			x:    1,
			want: 1, // Example: 1 = 1^1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.n, tt.x); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
