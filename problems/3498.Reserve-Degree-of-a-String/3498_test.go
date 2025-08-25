package leetcode

import (
	"testing"
)

func Test_reverseDegree(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abc",
			want: 148,
		},
		{
			name: "Example 2",
			s:    "zaza",
			want: 160,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseDegree(tt.s); got != tt.want {
				t.Errorf("reverseDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
