package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1493(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			},
			want: 5,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
		{
			name: "tc4 - all zeros",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "tc5 - single element",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "tc6 - mixed values",
			args: args{
				nums: []int{1, 0, 1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "tc7 - special case",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.args.nums)
			assert.Equal(t, tt.want, got, "longestSubarray(%v) should return %d, got %d", tt.args.nums, tt.want, got)
		})
	}
}
