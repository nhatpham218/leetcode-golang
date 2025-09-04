package leetcode

import "testing"

func TestFindClosest(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		z    int
		want int
	}{
		{
			name: "Example 1 - Person 1 closer",
			x:    2,
			y:    7,
			z:    4,
			want: 1,
		},
		{
			name: "Example 2 - Person 2 closer",
			x:    2,
			y:    5,
			z:    6,
			want: 2,
		},
		{
			name: "Example 3 - Same distance",
			x:    1,
			y:    5,
			z:    3,
			want: 0,
		},
		{
			name: "Person 1 at same position as Person 3",
			x:    5,
			y:    10,
			z:    5,
			want: 1,
		},
		{
			name: "Person 2 at same position as Person 3",
			x:    10,
			y:    5,
			z:    5,
			want: 2,
		},
		{
			name: "All at same position",
			x:    5,
			y:    5,
			z:    5,
			want: 0,
		},
		{
			name: "Person 3 between Person 1 and 2",
			x:    1,
			y:    9,
			z:    5,
			want: 0,
		},
		{
			name: "All at minimum constraint",
			x:    1,
			y:    1,
			z:    1,
			want: 0,
		},
		{
			name: "All at maximum constraint",
			x:    100,
			y:    100,
			z:    100,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findClosest(tt.x, tt.y, tt.z)
			if got != tt.want {
				t.Errorf("findClosest(%d, %d, %d) = %d, want %d", tt.x, tt.y, tt.z, got, tt.want)
			}
		})
	}
}
