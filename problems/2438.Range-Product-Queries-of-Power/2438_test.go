package leetcode

import (
	"testing"
)

func Test_productQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		queries [][]int
		want    []int
	}{
		{
			name: "tc1",
			n:    15,
			queries: [][]int{
				{0, 1},
				{2, 2},
				{0, 3},
			},
			want: []int{2, 4, 64},
		},
		{
			name: "tc2",
			n:    2,
			queries: [][]int{
				{0, 0},
			},
			want: []int{2},
		},
		{
			name: "tc3",
			n:    8,
			queries: [][]int{
				{0, 0},
			},
			want: []int{8},
		},
		{
			name: "tc4",
			n:    1,
			queries: [][]int{
				{0, 0},
			},
			want: []int{1},
		},
		{
			name: "tc5",
			n:    16,
			queries: [][]int{
				{0, 0},
			},
			want: []int{16},
		},
		{
			name: "tc6",
			n:    7,
			queries: [][]int{
				{0, 0},
				{1, 1},
				{2, 2},
				{0, 2},
			},
			want: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productQueries(tt.n, tt.queries); !compareSlices(got, tt.want) {
				t.Errorf("productQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to compare slices
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
