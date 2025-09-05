package leetcode

import "testing"

func TestMakeTheIntegerZero(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "Example 1",
			num1: 3,
			num2: -2,
			want: 3,
		},
		{
			name: "Example 2",
			num1: 5,
			num2: 7,
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeTheIntegerZero(tt.num1, tt.num2)
			if got != tt.want {
				t.Errorf("makeTheIntegerZero(%d, %d) = %d, want %d", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}
