package leetcode

import (
	"testing"
)

func Test_maxDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "tc1",
			s:    "NWSE",
			k:    1,
			want: 3,
		},
		{
			name: "tc2",
			s:    "NSWWEW",
			k:    3,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.s, tt.k); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
